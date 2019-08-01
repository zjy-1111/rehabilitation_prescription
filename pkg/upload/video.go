package upload

import (
	"log"
	"mime/multipart"
	"strings"

	"rehabilitation_prescription/pkg/file"
	"rehabilitation_prescription/pkg/logging"
	"rehabilitation_prescription/pkg/setting"
)

// CheckVideoExt check image file ext
func CheckVideoExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.VideoAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckVideoSize check image size
func CheckVideoSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.VideoMaxSize
}