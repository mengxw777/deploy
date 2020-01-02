package service

import (
	"deploy/app/models"
	"deploy/helper/command"
	"fmt"
	"strings"
)

const CommandTimeOut = 86400

const tarPath = "/tmp/deploy"

type CallbackFn func(string, *command.Task)

func Build(project models.Project, deploy models.Deploy, fn CallbackFn) {
	go func() {
		var task *command.Task
		var cmdList []string
		var currentBuildPath = fmt.Sprintf("%s/%d", tarPath, deploy.Id)

		// 	拉取代码、切换分支
		cmdList = append(cmdList, []string{
			"echo \"开始构建 Now is\" `date`",
			"echo \"Run user is\" `who am i`",
			fmt.Sprintf("/usr/bin/env git clone -q %s %s", project.GitAddr, currentBuildPath),
			fmt.Sprintf("cd %s && /usr/bin/env git checkout -q %s", currentBuildPath, project.Branch),
		}...)

		// 	自定义的构建命令
		for _, item := range strings.Split(project.BuildCmd, "\n") {
			cmdList = append(cmdList, []string{
				fmt.Sprintf("cd %s && %s", currentBuildPath, item),
			}...)
		}

		// 	打包
		zipName := fmt.Sprintf("%d.tar.gz", deploy.Id)
		cmdList = append(cmdList, []string{
			fmt.Sprintf("cd %s && tar --exclude='.git' -czvf %s *", currentBuildPath, tarPath+"/"+zipName),
			fmt.Sprintf("rm -rf %s", currentBuildPath),
			"echo \"构建完成 Now is\" `date`",
		}...)

		task = command.NewTask(cmdList, CommandTimeOut)
		task.Run()

		if fn != nil {
			fn(tarPath+"/"+zipName, task)
		}
	}()
}
