package imagersync

import (
	"time"

	"github.com/gin-gonic/gin"

	"novacloud-imagersync-service/internal/configVar"
)

// Get Url /image/rsync/me
func GetImageRsync(c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(200, gin.H{
		"username": username,
		"error":    "请使用 Post 请求进行 image rsync",
	})
}

// PostImageRsync: Post Url /image/rsync/
func PostImageRsync(c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	startTime := time.Now()
	requestJson := configVar.ImageRsyncRequestBody{}
	err := c.BindJSON(&requestJson)
	if err != nil {
		c.JSON(500, configVar.ImageRsyncResponseBody{
			RsyncData: configVar.ImageRsyncData{
				RsyncError: err.Error(),
				StartTime:  startTime,
			},
			RsyncConfig: configVar.ImageRsyncConfig{},
			Username:    username,
		})
		return
	}

	time.Sleep(time.Second * 2)

	endTime := time.Now()
	durationTimeSeconds := endTime.Sub(startTime).Seconds()
	imageRsyncResponseBody := configVar.ImageRsyncResponseBody{
		RsyncConfig: requestJson.RsyncConfig,
		Username:    username,
		RsyncData: configVar.ImageRsyncData{
			RsyncStatus:         "",
			RsyncError:          "",
			RsyncInfo:           "",
			StartTime:           startTime,
			EndTime:             endTime,
			DurationTimeSeconds: durationTimeSeconds,
		},
	}
	c.JSON(200, imageRsyncResponseBody)
}
