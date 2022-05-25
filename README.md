# GXS

#### 介绍
Go语言实验小说爬虫

#### 安装教程

> 提前安装好go语言

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
3.  全局使用
    ```sh
    go install
    export GOPATH=~/go
    export PATH=$PATH:$GOPATH/bin
    gxs [-b 网站域名] [-u 启始章节链接(不带域名)] -f 写入的文件 [-a] [-c 指定配置文件路径]
    ```

#### 配置文件

1. config.yaml 内容

```yaml
# config.yaml
base: https://www.qu-la.com # 小说网站的域名
purl: /booktxt/15394938116/61371543116.html # 开始页的链接去掉域名的部分
append_mode: false # 写文件为追加模式
encode: gbk # 网站的字符编码

query:
  title: ["h1.chaptername"] # 标题选择器
  content: ["div#txt"] #内容选择器
  next: ["a#pb_next"] #下一章选择器
  next_with_index: false #是否使用索引
  next_index: 0 #索引
```

2. 默认路径
    1. `./config.yaml`
    2. `~/.gxs.yaml`
    3. `/etc/gxs/config.yaml`