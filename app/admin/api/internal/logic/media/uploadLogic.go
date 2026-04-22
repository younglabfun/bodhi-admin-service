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
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
	_ "golang.org/x/image/webp" // 关键！注册 WebP 编码器
	"image"
	_ "image/jpeg" // 必须注册解码器
	_ "image/png"  // 必须注册解码器
	"io"
	"k8s.io/apimachinery/pkg/util/rand"
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

func (l *UploadLogic) Upload(r *http.Request) (*types.AffectedResp, error) {
	// 1. 解析 multipart form，限制上传大小为 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}

	// 2. 获取文件
	file, handler, err := r.FormFile("file") // "file" 是前端传参的 key
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 3. 保存文件到本地目录
	// 实际项目中建议使用唯一文件名（如 UUID）或时间戳
	dateDir := time.Now().Format("20060102") + "/"
	uploadDir := l.svcCtx.Config.AdminConf.UploadPath + dateDir
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		//http.Error(w, "Could not create directory", http.StatusInternalServerError)
		return nil, errors.New("could not create directory")
	}

	// 4. 创建目标文件
	//dstPath := filepath.Join(uploadDir, handler.Filename)
	dstPath, filename := generateUniqueFilename(uploadDir, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// 4. 拷贝文件数据
	_, err = io.Copy(dst, file)
	if err != nil {
		return nil, err
	}
	go func() {
		fmt.Println("process file--- ", filename)
		err = processImage(filename, uploadDir)
		if err != nil {
			fmt.Printf("process err: %v\n", err)
		}
	}()

	title := strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename))
	fileType := handler.Header.Get("Content-Type")
	path := dateDir + filename
	size := handler.Size
	var metaJson []byte
	width, height, err := getImageDimensions(dstPath)
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
		Type:     fileType,
		Path:     path,
		Size:     size,
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

// 处理上传图片的函数
func processImage(filename, outputDir string) error {
	// 1. 加载原图
	src, err := imaging.Open(outputDir + "/" + filename)
	if err != nil {
		return err
	}

	// 2. 生成列表缩略图
	small := imaging.Thumbnail(src, 100, 100, imaging.Lanczos)
	smallFilename := generateThumbnailFilename(filename, "small")
	smallImg, err := os.Create(outputDir + "/" + smallFilename)
	if err != nil {
		return err
	}
	defer smallImg.Close()
	err = webp.Encode(smallImg, small, &webp.Options{Lossless: false, Quality: 85})
	fmt.Println("p small--- ", smallFilename, err)

	// 3. 生成详情页预览图 (宽度 800px)
	medium := imaging.Resize(src, 800, 0, imaging.Lanczos)
	mediumFilename := generateThumbnailFilename(filename, "medium")
	mediumImg, err := os.Create(outputDir + "/" + mediumFilename)
	if err != nil {
		return err
	}
	defer mediumImg.Close()
	err = webp.Encode(mediumImg, medium, &webp.Options{Lossless: false, Quality: 95})
	fmt.Println("p medium--- ", mediumFilename, err)

	return err
}

// 生成缩略图名称
func generateThumbnailFilename(originalFilename, suffix string) string {
	ext := filepath.Ext(originalFilename)
	nameOnly := strings.TrimSuffix(originalFilename, ext)
	thumbnail := fmt.Sprintf("%s-%s%s", nameOnly, suffix, ".webp")
	return thumbnail
}

// 生成唯一文件名（如果存在则添加时间戳）
func generateUniqueFilename(uploadDir, originalFilename string) (string, string) {
	// 1. 获取后缀
	ext := filepath.Ext(originalFilename) // 获取扩展名，如 ".jpg"

	// 2. 生成唯一 ID (雪花算法)
	var (
		sonyFlake     *sonyflake.Sonyflake
		sonyMachineID uint16
		st            time.Time
		err           error
	)
	startTime := "2018-08-24" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	sonyMachineID = uint16(machineID)
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) { return sonyMachineID, nil },
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		panic("sonyflake not created")
	}

	id, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	// 3. 生成日期前缀
	//datePrefix := time.Now().Format("2006/01/02")

	// 4. 生成随机码
	randomSuffix := rand.String(6)

	// 5. 组合
	newFilename := fmt.Sprintf("%d_%s%s", id, randomSuffix, ext)
	return filepath.Join(uploadDir, newFilename), newFilename
}

func getImageDimensions(filePath string) (float64, float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// 只读取配置头，不加载图像像素数据，性能极佳
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	width := float64(config.Width)
	height := float64(config.Height)
	return width, height, nil
}
