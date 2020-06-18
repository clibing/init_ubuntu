package function

/**
 * 更新当前数据源
 */
import (
	log "github.com/sirupsen/logrus"
)

const SOURCE_LIST = `
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-security main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-proposed main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ OS_RELEASE-proposed main restricted universe multiverse
`

func setSourceList()  {
	log.Info("setting /etc/apt/source.list")


}



