package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")                                                                              // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("C:\\Users\\yida\\GolandProjects\\GoProjectDemo\\geek_time\\config\\config.yaml") // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.iam")                                                                     // 配置文件搜索路径，可以设置多个配置文件搜索路径
		viper.SetConfigName("config")
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			// 配置文件发生变更之后会调用的回调函数 fmt.Println("Config file changed:", e.Name)})// 配置文件名称（没有文件扩展名）
			fmt.Println("Config file changed:", e.Name)
		})
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}
