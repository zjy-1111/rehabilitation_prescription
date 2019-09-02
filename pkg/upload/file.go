package upload

import (
	"fmt"
	"os"
	"path"
	"rehabilitation_prescription/pkg/file"
	"strings"

	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/util"
)

// GetFileFullUrl get the full access path
func GetFileFullURL(name string) string {
	return "https://" + setting.AppSetting.BucketName + "." + setting.AppSetting.Endpoint + "/" + name
}

// GetFileName get image name
func GetFileName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetCoverName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ".jpg"
}

// CheckFile check if the file exists
func CheckFile(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
