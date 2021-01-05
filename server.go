package main

import (
	"checkpmk2/initialize"
	"github.com/labstack/echo"
	log "github.com/Sirupsen/logrus"
	"tester/controllers"
	"time"
	"os"
	"flag"
	"runtime"
	"github.com/labstack/echo/middleware"
)

func main()  {

	e :=echo.New()
	config:=initialize.Init{}
	init,err := initialize.InitConfig()
	config.Inits = init
	if err != nil  {
		log.Fatal(err)
	}
	goup :=e.Group("/v2/api")
	goup.POST("/tester",controllers.PMKCheking)
	now := time.Now()
	if runtime.GOOS =="windows"{
		logspath := config.PATHWINDOWS+config.FILELOGNAME+"."+now.Format("2006-01-02")+".log"
		os.Mkdir(config.PATHWINDOWS,0755)
		logs,err := os.OpenFile(logspath,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		Formmater :=new(log.TextFormatter)
		Formmater.TimestampFormat ="02-01-2006 15:04:05"
		Formmater.FullTimestamp = true
		log.SetFormatter(Formmater)
		if err != nil {
			log.Errorf("cannot open '"+logspath+"', (%s)", err.Error())
			flag.Usage()
			os.Exit(-1)
		}
		log.SetOutput(logs)
	}else{
		logspath := config.PATHLINUX+config.FILELOGNAME+"."+now.Format("2006-01-02")+".log"
		os.Mkdir(config.PATHLINUX,0755)
		logs,err := os.OpenFile(logspath,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		Formmater :=new(log.TextFormatter)
		Formmater.TimestampFormat ="02-01-2006 15:04:05"
		Formmater.FullTimestamp = true
		log.SetFormatter(Formmater)
		if err != nil {
			log.Errorf("cannot open '"+logspath+"', (%s)", err.Error())
			flag.Usage()
			os.Exit(-1)
		}
		log.SetOutput(logs)
	}
	e.Debug =true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	log.Info(e.Start(":"+config.Inits.PORT))
	e.Logger.Fatal(e.Start(":"+config.Inits.PORT))


}
