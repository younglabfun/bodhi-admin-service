package utils

import (
	"errors"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
	"image"
	"k8s.io/apimachinery/pkg/util/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var thumbnailDir = "/thumbnails"
var thumbnailConfig = map[string]int{
	"small":  200,
	"medium": 800,
}

func GetThumbnails(originalFilename string) map[string]string {
	var thumbnails = make(map[string]string)
	for size, _ := range thumbnailConfig {
		filename := GenerateThumbnailFilename(originalFilename, size)
		thumbnails[size] = filename
	}
	return thumbnails
}

func GetFilename(filePath string) (string, error) {
	if strings.Index(filePath, "/") < 0 {
		return "", errors.New("file path is error")
	}
	data := strings.Split(filePath, "/")
	return data[len(data)-1], nil
}

func RemoveFileAndThumbnails(filePath, uploadPath string) error {
	if strings.Index(filePath, "/") < 0 {
		return errors.New("file path is error")
	}
	filename, err := GetFilename(filePath)
	thumbnailPath := uploadPath + strings.ReplaceAll(filePath, "/"+filename, "") + thumbnailDir
	thumbnails := GetThumbnails(filename)
	for _, v := range thumbnails {
		t := thumbnailPath + "/" + v
		err := os.RemoveAll(t)
		if err != nil {
			logx.Errorf("remove file %s err %v\n", t, err)
		}
		//fmt.Printf("---- remove \n", t)
	}

	file := uploadPath + filePath
	err = os.RemoveAll(file)
	if err != nil {
		logx.Errorf("remove file %s err %v", file, err)
		return err
	}
	return nil
}

// 处理上传图片的函数
func DoProcessImage(filename, outputDir string) error {
	thumbnails := GetThumbnails(filename)
	if len(thumbnails) == 0 {
		return nil
	}

	// 1. 加载原图
	src, err := imaging.Open(outputDir + "/" + filename)
	if err != nil {
		return err
	}
	thumbnailPath := outputDir + thumbnailDir
	_, err = os.Stat(thumbnailPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(thumbnailPath, 0755)
		if err != nil {
			//http.Error(w, "Could not create directory", http.StatusInternalServerError)
			return errors.New("could not create directory")
		}
	}
	for key, thumbnailFile := range thumbnails {
		img := imaging.Resize(src, thumbnailConfig[key], 0, imaging.Lanczos) // 不裁切
		file, err := os.Create(thumbnailPath + "/" + thumbnailFile)
		if err != nil {
			return err
		}
		defer file.Close()
		_ = webp.Encode(file, img, &webp.Options{Lossless: false, Quality: 85})
	}

	//small := imaging.Thumbnail(src, 100, 100, imaging.Lanczos) // 裁切
	return err
}

func GetThumbnailFilename(filePath, suffix string) string {
	filename, _ := GetFilename(filePath)
	//fmt.Println("filename ", filePath, filename)
	if len(filename) == 0 {
		return ""
	}
	ext := filepath.Ext(filename)
	path := strings.ReplaceAll(filePath, "/"+filename, "")
	//fmt.Println("filePath --- ", path)
	nameOnly := strings.TrimSuffix(filename, ext)
	thumbnail := fmt.Sprintf("%s%s/%s-%s%s", path, thumbnailDir, nameOnly, suffix, ".webp")
	return thumbnail
}

// 生成缩略图名称
func GenerateThumbnailFilename(originalFilename, suffix string) string {
	ext := filepath.Ext(originalFilename)
	nameOnly := strings.TrimSuffix(originalFilename, ext)
	thumbnail := fmt.Sprintf("%s-%s%s", nameOnly, suffix, ".webp")
	return thumbnail
}

// 生成唯一文件名（如果存在则添加时间戳）
func GenerateUniqueFilename(uploadDir, originalFilename string) (string, string, error) {
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
		return "", "", err
	}

	sonyMachineID = uint16(machineID)
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) { return sonyMachineID, nil },
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		return "", "", errors.New("sonyflake not created")
	}

	id, err := sonyFlake.NextID()
	if err != nil {
		return "", "", err
	}

	// 3. 生成日期前缀
	//datePrefix := time.Now().Format("2006/01/02")

	// 4. 生成随机码
	randomSuffix := rand.String(6)

	// 5. 组合
	newFilename := fmt.Sprintf("%d_%s%s", id, randomSuffix, ext)
	return filepath.Join(uploadDir, newFilename), newFilename, nil
}

func GetImageDimensions(filePath string) (float64, float64, error) {
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
