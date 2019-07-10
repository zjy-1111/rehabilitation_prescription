package util

import "rehabilitation_prescription/pkg/setting"

// InitUtil Initialize the util
func InitUtil() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
