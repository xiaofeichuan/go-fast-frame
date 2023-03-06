## Go Fast Frame

## 项目启动
```
# 初始化依赖
go mod init

# 启动项目
go run main.go
```


## 目录结构

```
├── app                 (应用层)
│   ├── api             (API接口)
│   ├── dto             (数据传输对象)
│   └── service         (业务处理)
├── common              (公共文件)
│   ├── dto             (数据传输对象)
│   ├── tools           (初始化)
│   └── utils           (工具包)
├── config              (配置文件)
├── constant            (常量/枚举)
├── docs                (swagger文档目录)
├── global              (全局对象)
├── middleware          (中间件)
├── model               (数据模型)
├── router              (路由层)
├──	go.mod              (相关依赖)
├── go.sum              (相关依赖)
└── main.go             (入口文件)
```