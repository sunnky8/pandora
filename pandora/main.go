// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/ielepro/pandora"
	"github.com/ielepro/pandora/router/route"
	"github.com/ielepro/pandora/util/gopath"
)

var (
	helpFlag       bool
	pandoraIniFlag string
	versionFlag    bool
	configFile     *goconfig.ConfigFile
)

func init() {
	gin.SetMode(gin.DebugMode)

	flag.BoolVar(&helpFlag, "h", false, "This help")
	flag.StringVar(&pandoraIniFlag, "c", "", "Set configuration file `file`")
	flag.BoolVar(&versionFlag, "v", false, "Version number")

	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Printf("Usage: pandora [-c filename]\n\nOptions:\n")
	flag.PrintDefaults()
}

func initCfg() {
	var err error
	pandoraIni := findPandoraIniFile()
	configFile, err = goconfig.LoadConfigFile(pandoraIni)
	if err != nil {
		log.Fatalf("load config file failed, %s\n", err.Error())
	}
	outputInfo("Config Loaded", pandoraIni)
}

func configIntOrDefault(section, key string, useDefault int) int {
	val, err := configFile.Int(section, key)
	if err != nil {
		return useDefault
	}
	return val
}

func configOrDefault(section, key, useDefault string) string {
	val, err := configFile.GetValue(section, key)
	if err != nil {
		return useDefault
	}
	return val
}

func findPandoraIniFile() string {
	if pandoraIniFlag != "" {
		return pandoraIniFlag
	}
	currPath, _ := gopath.CurrentPath()
	parentPath, _ := gopath.CurrentParentPath()
	scanPath := []string{
		"/etc",
		currPath,
		fmt.Sprintf("%s/etc", currPath),
		fmt.Sprintf("%s/etc", parentPath),
	}

	for _, path := range scanPath {
		iniFile := path + "/pandora.ini"
		if gopath.Exists(iniFile) && gopath.IsFile(iniFile) {
			return iniFile
		}
	}

	return "./pandora.ini"
}

func outputInfo(tag string, value interface{}) {
	fmt.Printf("%-18s    %v\n", tag+":", value)
}

func welcome() {
	fmt.Println("潘多拉 自动化发布部署工具")
	outputInfo("Service", "pandora")
	outputInfo("Version", pandora.Version)
}

func main() {
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}
	if versionFlag {
		fmt.Printf("pandora/%s\n", pandora.Version)
		os.Exit(0)
	}

	welcome()

	initCfg()

	cfg := &pandora.Config{
		Serve: &pandora.ServeConfig{
			Addr:          configOrDefault("serve", "addr", "8868"),
			FeServeEnable: configIntOrDefault("serve", "fe_serve_enable", 1),
			ReadTimeout:   configIntOrDefault("serve", "read_timeout", 300),
			WriteTimeout:  configIntOrDefault("serve", "write_timeout", 300),
			IdleTimeout:   configIntOrDefault("serve", "idle_timeout", 300),
		},
		Db: &pandora.DbConfig{
			Unix:            configOrDefault("database", "unix", ""),
			Host:            configOrDefault("database", "host", ""),
			Port:            configIntOrDefault("database", "port", 3306),
			Charset:         "utf8mb4",
			User:            configOrDefault("database", "user", ""),
			Pass:            configOrDefault("database", "password", ""),
			DbName:          configOrDefault("database", "dbname", ""),
			MaxIdleConns:    configIntOrDefault("database", "max_idle_conns", 100),
			MaxOpenConns:    configIntOrDefault("database", "max_open_conns", 200),
			ConnMaxLifeTime: configIntOrDefault("database", "conn_max_life_time", 500),
		},
		Log: &pandora.LogConfig{
			Path: configOrDefault("log", "path", "stdout"),
		},
		Pandora: &pandora.PandoraConfig{
			LocalSpace:  configOrDefault("pandora", "local_space", "~/.pandora"),
			RemoteSpace: configOrDefault("pandora", "remote_space", "~/.pandora"),
			Cipher:      configOrDefault("pandora", "cipher_key", ""),
			AppHost:     configOrDefault("pandora", "app_host", ""),
		},
		Mail: &pandora.MailConfig{
			Enable: configIntOrDefault("mail", "enable", 0),
			Smtp:   configOrDefault("mail", "smtp_host", ""),
			Port:   configIntOrDefault("mail", "smtp_port", 465),
			User:   configOrDefault("mail", "smtp_user", ""),
			Pass:   configOrDefault("mail", "smtp_pass", ""),
		},
	}

	outputInfo("Log", cfg.Log.Path)
	outputInfo("Mail Enable", cfg.Mail.Enable)
	outputInfo("HTTP Service", cfg.Serve.Addr)

	pandora.App.Gin.Use(gin.Logger())

	if err := pandora.App.Init(cfg); err != nil {
		log.Fatal(err)
	}

	route.RegisterRoute()

	fmt.Println("Start Running...")
	if err := pandora.App.Start(); err != nil {
		log.Fatal(err)
	}
}
