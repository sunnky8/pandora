// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pandora

import (
    "errors"
    "encoding/base64"
    "fmt"
    "os"
    "io"

    "github.com/gin-gonic/gin"
    "github.com/ielepro/pandora/util/golog"
    "github.com/ielepro/pandora/util/gopath"
)

var (
    App	*pandora
)

const (
    Version = "v1.0.0"
)

func init() {
    App = newPandora()
}

type pandora struct {
    Gin             *gin.Engine
    DB              *DB
    Logger          *golog.Logger
    Mail            *SendMail
    LocalSpace      string
    LocalTmpSpace   string
    LocalTarSpace   string
    RemoteSpace     string
    CipherKey       []byte
    AppHost         string
    FeServeEnable   int
    config          *Config
}

func newPandora() *pandora {
    return &pandora{
        Gin: gin.New(),
    }
}

func (s *pandora) Init(cfg *Config) error {
    s.config = cfg

    if err := s.registerOrm(); err != nil {
        return err
    }
    s.registerMail()
    s.registerLog()

    if err := s.initEnv(); err != nil {
        return err
    }
    return nil
}

func (s *pandora) Start() error {
    return s.Gin.Run(s.config.Serve.Addr)
}

func  (s *pandora) registerOrm() error {
    s.DB = NewDatabase(s.config.Db)
    return s.DB.Open()
}

func (s *pandora) registerLog() {
    var loggerHandler io.Writer
    switch s.config.Log.Path {
    case "stdout":
        loggerHandler = os.Stdout
    case "stderr":
        loggerHandler = os.Stderr
    case "":
        loggerHandler = os.Stdout
    default:
        loggerHandler = golog.NewFileHandler(s.config.Log.Path)
    }
    s.Logger = golog.New(loggerHandler)
}

func (s *pandora) registerMail() {
    sendmail := &SendMail{
        Enable: s.config.Mail.Enable,
        Smtp: s.config.Mail.Smtp,
        Port: s.config.Mail.Port,
        User: s.config.Mail.User,
        Pass: s.config.Mail.Pass,
    }
    s.Mail = NewSendMail(sendmail)
}

func (s *pandora) initEnv() error {
    s.AppHost = s.config.Pandora.AppHost
    s.FeServeEnable = s.config.Serve.FeServeEnable
    s.LocalSpace = s.config.Pandora.LocalSpace
    s.LocalTmpSpace = s.LocalSpace + "/tmp"
    s.LocalTarSpace = s.LocalSpace + "/tar"

    if err := gopath.CreatePath(s.LocalSpace); err != nil {
        return err
    }
    if err := gopath.CreatePath(s.LocalTmpSpace); err != nil {
        return err
    }
    if err := gopath.CreatePath(s.LocalTarSpace); err != nil {
        return err
    }

    s.RemoteSpace = s.config.Pandora.RemoteSpace
    if s.config.Pandora.Cipher == "" {
        return errors.New("pandora config 'Cipher' not setting")
    }
    dec, err := base64.StdEncoding.DecodeString(s.config.Pandora.Cipher)
    if err != nil {
        return errors.New(fmt.Sprintf("decode Cipher failed, %s", err.Error()))
    }
    s.CipherKey = dec

    return nil
}
