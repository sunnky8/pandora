// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pandora

type (
    Config struct {
        Serve   *ServeConfig
        Db      *DbConfig
        Log     *LogConfig
        Pandora   *PandoraConfig
        Mail    *MailConfig
    }

    PandoraConfig struct {
        LocalSpace      string
        RemoteSpace     string
        Cipher          string
        AppHost         string
    }

    LogConfig struct {
        Path    string
    }

    MailConfig struct {
        Enable  int
        Smtp    string
        Port    int
        User    string
        Pass    string
    }

    ServeConfig struct {
        Addr            string
        FeServeEnable   int
        ReadTimeout     int
        WriteTimeout    int
        IdleTimeout     int
    }

    DbConfig struct {
        Unix            string
        Host            string
        Port            int
        Charset         string
        User            string
        Pass            string
        DbName          string
        TablePrefix     string
        MaxIdleConns    int
        MaxOpenConns    int
        ConnMaxLifeTime int
    }
)
