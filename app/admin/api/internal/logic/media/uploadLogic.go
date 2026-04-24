package media

import (
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	_ "golang.org/x/image/webp" // 关键！注册 WebP 编码器
	_ "image/jpeg"              // 必须注册解码器
	_ "image/png"               // 必须注册解码器
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// curl -X POST http: //admin.local/media/upload \
// -F "file=@/Users/yangjing/Projects/younglabs/bodhi-admin/IMG_0096.JPG"
func (l *UploadLogic) Upload(r *http.Request) (*types.AffectedResp, error) {
	// 1. 解析 multipart form
	maxMemory := int64(64 << 20)
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		return nil, err
	}

	// 2. 获取文件
	file, handler, err := r.FormFile("file") // "file" 是前端传参的 key
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 3. 创建目录，上传文件
	// 创建目录
	dateDir := time.Now().Format("20060102") + "/"
	uploadDir := l.svcCtx.Config.AdminConf.UploadPath + dateDir
	_, err = os.Stat(uploadDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return nil, errors.New("could not create directory")
		}
	}
	filename, err := doUploadFile(file, uploadDir, handler.Filename)
	if err != nil {
		return nil, err
	}

	go func() {
		fmt.Println("process file--- ", filename)
		err = utils.DoProcessImage(filename, uploadDir)
		if err != nil {
			fmt.Printf("process err: %v\n", err)
		}
	}()

	title := strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename))
	var metaJson []byte
	width, height, err := utils.GetImageDimensions(uploadDir + "/" + filename)
	fmt.Println("config--- ", width, height, err)
	if err == nil {
		aspectRatio := float64(width / height)
		meta := map[string]string{
			"aspect_ratio": utils.Float64ToStr(aspectRatio),
		}
		metaJson, _ = json.Marshal(meta)
	}
	resp, err := l.svcCtx.MediaRpc.InsertMedia(l.ctx, &admin.MediaReq{
		Title:    title,
		Filename: handler.Filename,
		Type:     handler.Header.Get("Content-Type"),
		Path:     dateDir + filename,
		Size:     handler.Size,
		Meta:     string(metaJson),
	})
	//fmt.Println("file--- ", path, fileType, size)
	//go func() {
	//	processImage(savePath, l.svcCtx.Config.AdminConf.UploadPath+"test.webp")
	//}()

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, err
}

func doUploadFile(file multipart.File, uploadPath, filename string) (string, error) {

	// 创建目标文件
	dstPath, uploadFile, err := utils.GenerateUniqueFilename(uploadPath, filename)
	if err != nil {
		return "", errors.New("could not create filename")
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// 4. 拷贝文件数据
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	return uploadFile, nil
}
