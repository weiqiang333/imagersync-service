package imagersync

import (
	"container/list"
	"fmt"

	"github.com/sirupsen/logrus"
	imageSyncClient "github.com/AliyunContainerService/image-syncer/pkg/client"
	"github.com/spf13/viper"
	"github.com/weiqiang333/imagersync-serviceimagersync-service/internal/imagersync"
)

// PushImageSync: 加载并运行镜像同步。
// parameter: <source_repo>:<dest_repo> map
func PushImageSync(source_repo string, dest_repo string)  {
	config := LoadAuthAndConfig(source_repo, dest_repo)
	client := imageSyncClient.Client{
		TaskList:                   list.New(),
		UrlPairList:                list.New(),
		FailedTaskList:             list.New(),
		FailedTaskGenerateList:     list.New(),
		Config:                     config,
		RoutineNum:                 viper.GetInt("rsyncserver.rsyncconfig.task.proc_num"),
		Retries:                    viper.GetInt("rsyncserver.rsyncconfig.task.retries"),
		Logger:                     imagersync.NewFileLogger(""),
		TaskListChan:               make(chan int, 1),
		UrlPairListChan:            make(chan int, 1),
		FailedTaskListChan:         make(chan int, 1),
		FailedTaskGenerateListChan: make(chan int, 1),
	}
	client.Run()
}

// LoadAuthConfig
func LoadAuthAndConfig(source_repo string, dest_repo string) *imageSyncClient.Config {
	authList := LoadAuth()
	var config = imageSyncClient.Config{
		AuthList: authList,
		ImageList: map[string]string{
			source_repo: dest_repo,
		},
	}
	fmt.Println(config)
	return &config
}

func LoadAuth() map[string]imageSyncClient.Auth {
	authMap := map[string]imageSyncClient.Auth{}
	// getViperConfigRegionTypeName: 这里提供了 dest_repo 的多个不同目标资源，这里按需更改
	getViperConfigRegionTypeName := viper.GetStringMap("rsyncserver.rsyncconfig.auth")
	if len(getViperConfigRegionTypeName) == 0 {
		fmt.Println("getViperConfigRegionTypeName: 并未提供镜像仓库的认证配置信息, 它确定对你没有认证影响嘛？")
	}

	for authName, _ := range getViperConfigRegionTypeName {
		authMap[viper.GetString(fmt.Sprintf("rsyncserver.rsyncconfig.auth.%s.registry", authName))] = imageSyncClient.Auth{
			Username: viper.GetString(fmt.Sprintf("rsyncserver.rsyncconfig.auth.%s.username", authName)),
			Password: viper.GetString(fmt.Sprintf("rsyncserver.rsyncconfig.auth.%s.password", authName)),
			Insecure: false,
		}
	}
	fmt.Println(authMap)
	return authMap
}

