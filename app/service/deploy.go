package service

import (
	"deploy/app/models"
	"deploy/helper/command"
	"fmt"
	"strings"
)

func Deploy(deploy models.Deploy, project models.Project, serverList []models.Server, fn CallbackFn) {
	go func() {
		var task *command.Task
		var cmdList []string
		for _, server := range serverList {
			ip := server.Ip

			// 	部署前置命令
			cmdList = append(cmdList, []string{
				"echo \"开始部署 Now is\" `date`",
				"echo \"Run user is\" `who am i`",
			}...)

			for _, item := range strings.Split(project.DeployBeforeCmd, "\n") {
				cmdList = append(cmdList, []string{
					fmt.Sprintf("ssh %s@%s %s", project.DeployUser, ip, item),
				}...)
			}

			cmdList = append(cmdList, []string{
				fmt.Sprintf("ssh %s@%s 'mkdir -p %s'", project.DeployUser, ip, project.DeployPath),
				fmt.Sprintf("scp -P %d %s %s@%s:%s", server.Port, deploy.BuildPath, project.DeployUser, ip, tarPath),
				fmt.Sprintf("ssh %s@%s 'tar -zxvf %s -C %s'", project.DeployUser, ip, deploy.BuildPath, project.DeployPath),
			}...)

			// 	部署后置命令
			for _, item := range strings.Split(project.DeployAfterCmd, "\n") {
				cmdList = append(cmdList, []string{
					fmt.Sprintf("ssh %s@%s %s", project.DeployUser, ip, item),
				}...)
			}

			cmdList = append(cmdList, []string{
				"echo \"部署完成 Now is\" `date`",
			}...)
		}

		task = command.NewTask(cmdList, CommandTimeOut)
		task.Run()

		if fn != nil {
			fn("", task)
		}
	}()
}
