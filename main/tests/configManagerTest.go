package tests

import "go_demo/configManager"

func  ReadConfig(){
	configManager.GetRedisConfig()
	configManager.GetMysqlConfig()
}
