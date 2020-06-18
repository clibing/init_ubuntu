/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

const BUFFERSIZE = 4096

var sourceFile string
var targetFile string

//focal
const debainSourceList = `# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-security main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-proposed main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ VERSION-proposed main restricted universe multiverse`

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info(sourceFile, targetFile)
		log.Info(args)
		copy(sourceFile, targetFile)
		if isLinux() {
			ubuntuDebian(cmd, args, sourceFile)
		}

	},
}

func darwin(cmd *cobra.Command, args []string, sourceFile string, targetFile string) (e error) {
	// 源文件
	source, err := os.OpenFile(sourceFile, os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Error(err)
		Error(cmd, args, err, true)
	}
	codeName := ""
	if isLinux() {
		codeName, err = ExecuteCommand("lsb_release", "-cs")
		if err != nil {
			Error(cmd, args, err, true)
		}
	}
	if isDarwin() {
		codeName, err = ExecuteCommand("uname", "-s")
		if err != nil {
			Error(cmd, args, err, true)
		}
	}
	codeName = strings.Replace(codeName, "\n", "", -1)
	content := strings.Replace(debainSourceList, "VERSION", codeName, -1)
	_, err = source.WriteString(content)
	if err != nil {
		Error(cmd, args, err, true)
	}
	defer source.Close()

	if isLinux() {
		output, err := ExecuteCommand("apt", "update")

		if err != nil {
			log.Error(output)
			Error(cmd, args, err, true)
		}
	}
	return
}

func ubuntuDebian(cmd *cobra.Command, args []string, sourceFile string) (e error) {
	// 源文件
	source, err := os.OpenFile(sourceFile, os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Error(err)
		Error(cmd, args, err, true)
	}

	codeName, err := ExecuteCommand("lsb_release", "-cs")
	if err != nil {
		Error(cmd, args, err, true)
	}
	codeName = strings.Replace(codeName, "\n", "", -1)

	content := strings.Replace(debainSourceList, "VERSION", codeName, -1)
	_, err = source.WriteString(content)
	if err != nil {
		Error(cmd, args, err, true)
	}
	defer source.Close()

	output, err := ExecuteCommand("apt", "update")
	if err != nil {
		log.Error(output)
		Error(cmd, args, err, true)
	}
	return
}

func copy(sourceFile string, targetFile string) (e error) {
	buf := make([]byte, BUFFERSIZE)
	// 源文件
	source, err := os.Open(sourceFile)
	if err != nil {
		e = err
		return
	}
	defer source.Close()

	// 目标文件
	destination, err := os.Create(targetFile)
	if err != nil {
		e = err
		return
	}
	defer destination.Close()

	// 循环写入数据
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}
func init() {
	updateCmd.Flags().StringVarP(&sourceFile, "source", "s", "", "copy source file path")
	updateCmd.Flags().StringVarP(&targetFile, "target", "t", "", "copy target file path")

	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
