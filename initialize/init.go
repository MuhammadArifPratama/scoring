package initialize

import (
"github.com/spf13/viper"
log "github.com/Sirupsen/logrus"

)
type Inits struct {
	MODE string
	PORT string
	Dbmysql struct{
		Driver string
		Username string
		Password string
		DatabaseName string
		Host string
	}
	DBTableName string
	FILELOGNAME string
	PATHLINUX string
	PATHWINDOWS string
}

func InitConfig() (Inits, error) {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("tsconfig")

	err :=viper.ReadInConfig()

	if err !=nil{
		log.Error("Confing error",err)
		return Inits{},err
	}
	viper.SetDefault("PORT","8080")

	var constants Inits
	err = viper.Unmarshal(&constants)
	return constants, err

}

