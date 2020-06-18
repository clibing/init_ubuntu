package function

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

func Cmd(cmd string, shell bool) (value []byte, e error) {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error("some error found")
			e = err
			return
		}
		value = out
		return
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			log.Error("some error found")
			e = err
			return
		}
		value = out
		return
	}
}

// 检查是否为linux system
func linuxUnix() {
	if runtime.GOOS == "windows" {
		panic("暂不支持windows")
	}
}

func Release(codeName bool)(value string) {
	linuxUnix()
	param := "a"
	if codeName {
		param = "c"
	}
	info, err := Cmd(fmt.Sprintf("lsb_release -%s", param), true)
	if err != nil {
		panic("暂无lsb_release命令")
	}
	log.Info(info)
	if codeName {
		value = string(info)
	}
	return
}
