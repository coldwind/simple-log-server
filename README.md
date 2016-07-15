# cw_log_server
> cw_log_server是一个udp日志服务器
> 默认存储日志在 bin/data/ 下
> 服务器以小时为单位创建日志记录的文件夹，每个文件夹内最多25个.data文件

## DATA命名规则
- 目录文件 年月日时 如：2016010101
- 日志文件 年-月-日-时＋随机数.data 如：2016-01-01-0130.data

## 配置文件
- bin/base.ini
- path 修改日志文件的存放目录
- ip 服务器的ip地址
- port 服务器开放的端口

## 安装&使用
go install cw_log_server
./cw_log_server
测试：nc -u 127.0.0.1 19851 (ip port以配置为准)
