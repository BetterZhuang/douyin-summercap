/**
    @author: zzg
    @date: 2022/5/27 22:05
    @dir_path: settings
    @note:
**/

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Conf 全局变量，用来保存程序所有的配置信息
var Conf = new(AppConfig) //指针，同一地址，同步更新

type AppConfig struct {
	Name         string `mapstructure:"name"` //注意使用mapstructure 改成.json也不用更改
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	//*MySQLConfig `mapstructure:"mysql"`
	//*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

//type MySQLConfig struct {
//	Host         string `mapstructure:"host"`
//	User         string `mapstructure:"user"`
//	Password     string `mapstructure:"password"`
//	DbName       string `mapstructure:"dbname"`
//	Port         int    `mapstructure:"port"`
//	MaxOpenConns int    `mapstructure:"max_open_conns"`
//	MaxIdleConns int    `mapstructure:"max_idle_conns"`
//}

//type RedisConfig struct {
//	Host     string `mapstructure:"host"`
//	Password string `mapstructure:"password"`
//	Port     int    `mapstructure:"port"`
//	DB       int    `mapstructure:"db"`
//	PoolSize int    `mapstructure:"pool_size"`
//}

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") //指定配置文件
	err = viper.ReadInConfig() //读取配置信息
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInitConfig() failed, err:%v\n", err)
		return
		//panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	//将读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	viper.WatchConfig()                            //实时监控配置文件变动情况
	viper.OnConfigChange(func(in fsnotify.Event) { //钩子，回调函数
		fmt.Println("配置文件修改！！！")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}