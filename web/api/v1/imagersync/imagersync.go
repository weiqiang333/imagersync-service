package imagersync

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weiqiang333/imagersync-service/internal/configVar"
	"github.com/weiqiang333/imagersync-service/internal/imagersync"
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
	logger := imagersync.NewStdoutLogger()
	err := c.BindJSON(&requestJson)
	if err != nil {
		logger.Errorf("failed in PostImageRsync: requestJson is not BindJSON, error: %s", err.Error())
		c.JSON(500, configVar.ImageRsyncResponseBody{
			RsyncConfig: configVar.ImageRsyncConfig{},
			Username:    username,
			RsyncData: configVar.ImageRsyncData{
				RsyncStatus: "unusual",
				RsyncError:  err.Error(),
				StartTime:   startTime,
			},
		})
		return
	}

	logger.Infof("in PostImageRsync 开始同步: user: %s, RsyncConfig: %s", username, requestJson.RsyncConfig)
	rsyncStatus, processData, err := imagersync.PushImageSync(requestJson.RsyncConfig.Source, requestJson.RsyncConfig.Target)

	endTime := time.Now()
	durationTimeSeconds := endTime.Sub(startTime).Seconds()

	imageRsyncResponseBody := configVar.ImageRsyncResponseBody{
		RsyncConfig: requestJson.RsyncConfig,
		Username:    username,
		RsyncData: configVar.ImageRsyncData{
			RsyncStatus:         rsyncStatus,
			RsyncError:          err,
			RsyncInfo:           processData,
			StartTime:           startTime,
			EndTime:             endTime,
			DurationTimeSeconds: durationTimeSeconds,
		},
	}

	if err != nil {
		logger.Errorf("failed in PostImageRsync: 同步异常, user: %s, RsyncConfig: %s, error: %s",
			username, requestJson.RsyncConfig, err.Error())
		c.JSON(500, imageRsyncResponseBody)
		return
	}

	c.JSON(200, imageRsyncResponseBody)
}
