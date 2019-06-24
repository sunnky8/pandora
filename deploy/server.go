// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package deploy

import (
	"github.com/ielepro/pandora/k8s"
	"github.com/ielepro/pandora/util/command"
)

const (
	COMMAND_TIMEOUT = 3600
)

type Server struct {
	ID            int
	Addr          string
	User          string
	Port          int
	PreCmd        string
	PostCmd       string
	Key           string
	PackFile      string
	DeployTmpPath string
	DeployPath    string
	task          *command.Task
	result        *ServerResult
}

type ServerResult struct {
	ID         int
	TaskResult []*command.TaskResult
	Status     int
	Error      error
}

func NewServer(srv *Server) {
	srv.result = &ServerResult{
		ID:     srv.ID,
		Status: STATUS_INIT,
	}
	srv.task = command.NewTask(
		srv.deployCmd(),
		COMMAND_TIMEOUT,
	)
}

func (srv *Server) Deploy() {
	srv.result.Status = STATUS_ING
	//srv.task.Run()

	deployment := k8s.Deployment{}
	if err := deployment.Update("nginx-deployment"); err != nil {
		srv.result.Status = STATUS_FAILED
	}

	if err := srv.task.GetError(); err != nil {
		srv.result.Error = err
		srv.result.Status = STATUS_FAILED
	} else {
		srv.result.Status = STATUS_DONE
	}
}

func (srv *Server) Terminate() {
	if srv.result.Status == STATUS_ING {
		srv.task.Terminate()
	}
}

func (srv *Server) Result() *ServerResult {
	srv.result.TaskResult = srv.task.Result()
	return srv.result
}

func (srv *Server) deployCmd() []string {
	var cmds []string
	return cmds
}
