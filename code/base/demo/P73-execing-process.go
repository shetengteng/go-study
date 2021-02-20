package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 切换到执行路径 如 ls路径下
	path, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	// 命令语句，第一个是命令，后面的是该命令的参数
	args := []string{"ls", "-a", "-l", "-h"}
	// 获取环境参数
	env := os.Environ()
	// 执行对应的命令，执行必须放在main语句最后面
	// 执行成功后，当前的进程id会给执行的命令的进程使用
	execErr := syscall.Exec(path, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
