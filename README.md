# GXS

#### 介绍
Go语言实验小说爬虫

#### 软件架构
软件架构说明


#### 安装教程

1.  `git clone https://gitee.com/uomer/gxs`
2.  `cd gxs`
3.  `go build`
4.  `cp config.yaml.backup config.yaml`
5.  全局安装使用的话，需要选处理配置文件或命令行中指定配置文件
    * 默认配置文件有以下路径
        1. ~/.gxs.yaml
        2. /etc/gxs/config.yaml
    * 编译安装 `go install`

#### 使用说明

1.  修改 config.yaml
2.  `./gxs [-b 网站域名] [-u 启始章节链接(不带域名)] -f 写入的文件 [-a] [-c 指定配置文件路径]`
    - [] 内为可选项
    - -a 为追加模式写入
    - 还可以用 -c 选项，指定要使用的配置文件
