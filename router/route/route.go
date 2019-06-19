// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
    "github.com/ielepro/pandora/router/k8s"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/ielepro/pandora/util/gopath"
    "github.com/ielepro/pandora/util/gostring"
    reqApi "github.com/ielepro/pandora/router/route/api"
    "github.com/ielepro/pandora"
    "github.com/ielepro/pandora/router/user"
    "github.com/ielepro/pandora/router/server"
    "github.com/ielepro/pandora/router/project"
    "github.com/ielepro/pandora/router/deploy"
    "github.com/ielepro/pandora/router/middleware"
)

func RegisterRoute() {
    api := pandora.App.Gin.Group("/api", middleware.ApiPriv())
    {
        api.POST(reqApi.LOGIN, user.Login)
        api.POST(reqApi.LOGOUT, user.Logout)
        api.GET(reqApi.LOGIN_STATUS, user.LoginStatus)

        api.POST(reqApi.MY_USER_SETTING, user.MyUserSetting)
        api.POST(reqApi.MY_USER_PASSWORD, user.MyUserPassword)

        api.POST(reqApi.SERVER_GROUP_ADD, server.GroupAdd)
        api.GET(reqApi.SERVER_GROUP_LIST, server.GroupList)
        api.POST(reqApi.SERVER_GROUP_DELETE, server.GroupDelete)
        api.GET(reqApi.SERVER_GROUP_DETAIL, server.GroupDetail)
        api.POST(reqApi.SERVER_GROUP_UPDATE, server.GroupUpdate)
        api.POST(reqApi.SERVER_ADD, server.ServerAdd)
        api.POST(reqApi.SERVER_UPDATE, server.ServerUpdate)
        api.GET(reqApi.SERVER_LIST, server.ServerList)
        api.POST(reqApi.SERVER_DELETE, server.ServerDelete)
        api.GET(reqApi.SERVER_DETAIL, server.ServerDetail)

        api.GET(reqApi.USER_ROLE_PRIV_LIST, user.RolePrivList)
        api.POST(reqApi.USER_ROLE_ADD, user.RoleAdd)
        api.POST(reqApi.USER_ROLE_UPDATE, user.RoleUpdate)
        api.GET(reqApi.USER_ROLE_LIST, user.RoleList)
        api.GET(reqApi.USER_ROLE_DETAIL, user.RoleDetail)
        api.POST(reqApi.USER_ROLE_DELETE, user.RoleDelete)
        api.POST(reqApi.USER_ADD, user.UserAdd)
        api.POST(reqApi.USER_UPDATE, user.UserUpdate)
        api.GET(reqApi.USER_LIST, user.UserList)
        api.GET(reqApi.USER_EXISTS, user.UserExists)
        api.GET(reqApi.USER_DETAIL, user.UserDetail)
        api.POST(reqApi.USER_DELETE, user.UserDelete)

        api.POST(reqApi.PROJECT_SPACE_ADD, project.SpaceAdd)
        api.POST(reqApi.PROJECT_SPACE_UPDATE, project.SpaceUpdate)
        api.GET(reqApi.PROJECT_SPACE_LIST, project.SpaceList)
        api.GET(reqApi.PROJECT_SPACE_DETAIL, project.SpaceDetail)
        api.POST(reqApi.PROJECT_SPACE_DELETE, project.SpaceDelete)
        api.GET(reqApi.PROJECT_MEMBER_SEARCH, project.MemberSearch)
        api.POST(reqApi.PROJECT_MEMBER_ADD, project.MemberAdd)
        api.GET(reqApi.PROJECT_MEMBER_LIST, project.MemberList)
        api.POST(reqApi.PROJECT_MEMBER_REMOVE, project.MemberRemove)
        api.POST(reqApi.PROJECT_ADD, project.ProjectAdd)
        api.POST(reqApi.PROJECT_UPDATE, project.ProjectUpdate)
        api.GET(reqApi.PROJECT_LIST, project.ProjectList)
        api.POST(reqApi.PROJECT_SWITCHSTATUS, project.ProjectSwitchStatus)
        api.GET(reqApi.PROJECT_DETAIL, project.ProjectDetail)
        api.POST(reqApi.PROJECT_DELETE, project.ProjectDelete)
        api.POST(reqApi.PROJECT_BUILDSCRIPT, project.ProjectBuildScript)
        api.POST(reqApi.PROJECT_HOOKSCRIPT, project.ProjectHookScript)

        api.GET(reqApi.DEPLOY_APPLY_PROJECT_DETAIL, deploy.ApplyProjectDetail)
        api.POST(reqApi.DEPLOY_APPLY_SUBMIT, deploy.ApplySubmit)
        api.GET(reqApi.DEPLOY_APPLY_PROJECT_ALL, deploy.ApplyProjectAll)
        api.GET(reqApi.DEPLOY_APPLY_LIST, deploy.ApplyList)
        api.GET(reqApi.DEPLOY_APPLY_DETAIL, deploy.ApplyDetail)
        api.POST(reqApi.DEPLOY_APPLY_AUDIT, deploy.ApplyAudit)
        api.POST(reqApi.DEPLOY_APPLY_UPDATE, deploy.ApplyUpdate)
        api.POST(reqApi.DEPLOY_APPLY_DROP, deploy.ApplyDrop)
        api.POST(reqApi.DEPLOY_BUILD_START, deploy.BuildStart)
        api.GET(reqApi.DEPLOY_BUILD_STATUS, deploy.BuildStatus)
        api.POST(reqApi.DEPLOY_BUILD_STOP, deploy.BuildStop)
        api.POST(reqApi.DEPLOY_DEPLOY_START, deploy.DeployStart)
        api.GET(reqApi.DEPLOY_DEPLOY_STATUS, deploy.DeployStatus)
        api.POST(reqApi.DEPLOY_DEPLOY_STOP, deploy.DeployStop)
        api.GET(reqApi.DEPLOY_APPLY_ROLLBACK, deploy.ApplyRollbackList)
        api.POST(reqApi.DEPLOY_DEPLOY_ROLLBACK, deploy.DeployRollback)

        api.GET(reqApi.K8S_DEPLOYMENT_LIST, k8s.DeploymentList)
        api.GET(reqApi.K8S_RESOURCES, k8s.Resources)
    }

    if pandora.App.FeServeEnable == 1 {
        RegisterFeResource()
    }
}

func RegisterFeResource() {
    parentPath, err := gopath.CurrentParentPath()
    if err != nil {
        pandora.App.Logger.Error("get current path failed, err[%s]", err.Error())
        return
    }
    indexFile := gostring.JoinStrings(parentPath, "/public/index.html")
    staticPath := gostring.JoinStrings(parentPath, "/public")

    pandora.App.Gin.StaticFile("/", indexFile)
    pandora.App.Gin.Static("/static", staticPath)

    pandora.App.Gin.LoadHTMLFiles(indexFile)
    pandora.App.Gin.NoRoute(func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })
}