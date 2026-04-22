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
	"github.com/google/uuid"
	_ "golang.org/x/image/webp" // 关键！注册 WebP 编码器
	"image"
	_ "image/jpeg" // 必须注册解码器
	_ "image/png"  // 必须注册解码器
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
		fmt.Println("file--- ", dstPath)
		_ = processImage(file, filename, uploadDir)
	}()

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
		Title:    "",
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
func processImage(file multipart.File, filename string, outputDir string) error {
	// 1. 加载原图
	//src, err := imaging.Open(filePath)
	//if err != nil {
	//	return err
	//}
	// file 是 multipart.File 类型
	src, _, err := image.Decode(file) // 使用标准库的 image.Decode 获取 image.Image
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

	// 3. 生成详情页预览图 (宽度 800px)
	medium := imaging.Resize(src, 800, 0, imaging.Lanczos)
	mediumFilename := generateThumbnailFilename(filename, "medium")
	err = imaging.Save(medium, outputDir+"/"+mediumFilename)
	fmt.Println("p medium--- ", err)

	// 使用 webp 编码器直接编码
	//err = imaging.Encode(out, medium, imaging.Format(webp.WebP))
	//err = webp.Encode(out, src, &webp.Options{Lossless: false, Quality: 85})

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
	dstPath := filepath.Join(uploadDir, originalFilename)

	// 检查文件是否存在
	if _, err := os.Stat(dstPath); err == nil {
		// 文件存在，处理逻辑：文件名 + 时间戳 + 后缀
		ext := filepath.Ext(originalFilename)                 // 获取扩展名，如 ".jpg"
		nameOnly := strings.TrimSuffix(originalFilename, ext) // 获取文件名主体，如 "photo"
		newFilename := fmt.Sprintf("%s-%s%s", nameOnly, uuid.New().String(), ext)

		return filepath.Join(uploadDir, newFilename), newFilename
	}

	// 文件不存在，直接返回原路径
	return dstPath, originalFilename
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
