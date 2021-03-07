package configManager

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	MysqlConfig Mysql `mapstructure:"mysql"`
	RedisConfig Redis `mapstructure:"redis"`
}

type Mysql struct {
	ReadDB struct {
		DriveName string `mapstructure:"drivername"`
		Notes     string `mapstructure:"notes"`
		Host      string `mapstructure:"host"`
		Port      string `mapstructure:"port"`
		DbName    string `mapstructure:"dbname"`
		User      string `mapstructure:"user"`
		Password  string `mapstructure:"password"`
	} `mapstructure:"readDB"`
	WriteDB struct {
		DriveName string `mapstructure:"drivername"`
		Notes     string `mapstructure:"notes"`
		Host      string `mapstructure:"host"`
		Port      string `mapstructure:"port"`
		DbName    string `mapstructure:"dbname"`
		User      string `mapstructure:"user"`
		Password  string `mapstructure:"password"`
	} `mapstructure:"writeDB"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var Conf Config

func readConfig() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 将读取的配置信息保存至全局变量dbConfig
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	fmt.Println(Conf.MysqlConfig)
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量dbConfig
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
}
func getConfig() {
	if Conf == (Config{}) {
		readConfig()
	}
}

func GetMysqlConfig() Mysql {
	getConfig()
	fmt.Println(Conf.MysqlConfig.WriteDB)
	fmt.Println(Conf.MysqlConfig.ReadDB)
	return Conf.MysqlConfig
}

func GetRedisConfig() Redis {
	getConfig()
	fmt.Print(Conf.MysqlConfig.ReadDB.User)
	fmt.Print(Conf)
	return Conf.RedisConfig
}
