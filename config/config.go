package config

import "github.com/spf13/viper"

type config struct{
	viper *viper.Viper
}
type tables struct{
	viper *viper.Viper
}

var Confs *config
var Tables *tables

//初始化表
func init(){
	Confs = &config{
		getConf(),
	}
	Tables = &tables{
		getTables(),
	}
}

//得到config配置文件
func getConf() *viper.Viper{
	v:=viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	v.ReadInConfig()
	return v
}

//得到 tablename 的配置文件
func getTables() *viper.Viper{
	v:=viper.New()
	v.SetConfigName("tableNames")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	v.ReadInConfig()
	return v
}

//得到相应的配置值
func (c *config)GetString(key string)string{
	return c.viper.GetString(key)
}
func (t *tables)GetString(key string)string{
	return t.viper.GetString(key)
}