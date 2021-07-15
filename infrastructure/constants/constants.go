package constants

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	ACCESS_KEY string
	SECRET_KEY string
	ACCESS_TOKEN string
)

type orderBy struct{
	RELEVANT string
	FEATURED string
	LATEST string
	OLDEST string
	POSITION string
	POPULAR string
}

var PhotosOrderBy = orderBy{
	LATEST: "latest",
	OLDEST: "oldest",
	POPULAR: "popular",
}

type ContentFilter string
const (
	LOW ContentFilter = "low"
	HIGH ContentFilter = "high"
)

type Color string
const (
	BLACK_AND_WHITE Color = "black_and_white"
	BLACK Color = "black"
	WHITE Color = "white"
	YELLOW Color = "yellow"
	ORANGE Color = "orange"
	RED Color = "red"
	PURPLE Color = "purple"
	MAGENTA Color = "magenta"
	GREEN Color = "green"
	TEAL Color = "teal"
	BLUE Color = "blue"
)

type Orientation string
const (
	LANDSCAPE Orientation = "landscape"
	PORTRAIT Orientation = "portrait"
	SQUARISH Orientation = "squarish"
)

type Bool string
const (
	TRUE Bool = "true"
	FALSE Bool = "false"
)

func init()  {
	//viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	ACCESS_KEY = viper.GetString("ACCESS_KEY")
	ACCESS_TOKEN = viper.GetString("ACCESS_TOKEN")
	SECRET_KEY = viper.GetString("SECRET_KEY")

}
