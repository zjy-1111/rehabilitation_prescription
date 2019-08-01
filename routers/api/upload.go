package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rehabilitation_prescription/pkg/ali_oss"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/logging"
	"rehabilitation_prescription/pkg/upload"
)

func UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if file == nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetFileName(image.Filename)
	uploadPath := "avatar/" + imageName

	err = ali_oss.Bucket.PutObject(uploadPath, file)
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"url": upload.GetFileFullURL(uploadPath),
	})
}

func UploadVideo(c *gin.Context) {
	file, video, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if file == nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	videoName := upload.GetFileName(video.Filename)
	uploadPath := "video/" + videoName

	err = ali_oss.Bucket.PutObject(uploadPath, file)
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	videoURL := upload.GetFileFullURL(uploadPath)

	// coverURL
	//coverName := upload.GetCoverName(videoName)
	//coverUploadPath := "video/" + coverName
	//
	//img, err := util.GetVideoMoment(videoURL, 0.0)
	//if err != nil {
	//	logging.Warn(err)
	//	app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
	//	return
	//}

	//imgFile, err := os.Create(coverName)
	//if err != nil {
	//	logging.Warn(err)
	//	app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
	//	return
	//}
	//jpeg.Encode(imgFile, img, nil)
	//
	//
	//err = ali_oss.Bucket.PutObject(coverUploadPath, imgFile)
	//if err != nil {
	//	logging.Warn(err)
	//	app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
	//	return
	//}
	//coverURL := upload.GetFileFullURL(uploadPath)
	//
	//// duration
	//duration, err := util.GetVideoDuration(videoURL)
	//if err != nil {
	//	logging.Warn(err)
	//	app.Response(c, http.StatusInternalServerError, e.ERROR, nil)
	//	return
	//}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"video_url": videoURL,
		//"cover_url": coverURL,
		//"duration": strconv.FormatFloat(duration, 'f', 2, 64),
	})
}

