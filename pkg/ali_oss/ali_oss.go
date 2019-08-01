package ali_oss

import (
	"os"
	"rehabilitation_prescription/pkg/logging"
	"rehabilitation_prescription/pkg/setting"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Client *oss.Client
var Bucket *oss.Bucket

func InitOssBucket() error {
	var err error
	Client, err = oss.New(
		setting.AppSetting.Endpoint,
		setting.AppSetting.OssAccessKeyId,
		setting.AppSetting.OssAccessKeySecret,
	)
	if err != nil {
		logging.Error(err)
		return err
	}

	Bucket, err = Client.Bucket("study-fs")
	if err != nil {
		logging.Error(err)
		os.Exit(-1)
	}

	return nil
}
