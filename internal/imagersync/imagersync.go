package imagersync

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	imageSyncClient "github.com/weiqiang333/image-syncer/pkg/client"
)

// PushImageSync: 加载并运行镜像同步。
// parameter: <source_repo>:<dest_repo> map
// return: rsync status, rsync process Data, error
func PushImageSync(source_repo string, dest_repo string) (string, string, error) {
	config := LoadAuthAndConfig(source_repo, dest_repo)
	log, processData := NewStdoutAndBufferLogger()
	client := imageSyncClient.Client{
		TaskList:                   list.New(),
		UrlPairList:                list.New(),
		FailedTaskList:             list.New(),
		FailedTaskGenerateList:     list.New(),
		Config:                     config,
		RoutineNum:                 viper.GetInt("rsyncserver.rsyncconfig.task.proc_num"),
		Retries:                    viper.GetInt("rsyncserver.rsyncconfig.task.retries"),
		Logger:                     log,
		TaskListChan:               make(chan int, 1),
		UrlPairListChan:            make(chan int, 1),
		FailedTaskListChan:         make(chan int, 1),
		FailedTaskGenerateListChan: make(chan int, 1),
	}
	client.Run()

	rsyncStatus := analysisProcessData(processData.String())

	return rsyncStatus, processData.String(), nil
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
	return &config
}

//LoadAuth:
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
			Insecure: viper.GetBool(fmt.Sprintf("rsyncserver.rsyncconfig.auth.%s.insecure", authName)),
		}
	}
	return authMap
}

// analysisProcessData: 分析过程数据
func analysisProcessData(processData string) string {
	if strings.Contains(processData, "Synchronization successfully") {
		return "successfully"
	}
	return "failed"
}
