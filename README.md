## 运行

    go run main.go

## 目录结构

    /
    ├── app          // iris设置及路由
    ├── config       // 配置
    ├── controllers  // 控制层
    ├── middleware   // iris中间件
    ├── model        // 实体
    ├── public       // 静态目录
    ├── repository   // 仓库
    ├── runtime      // logs等运行数据
    ├── services     // 服务层
    ├── util         // 工具
    └── views        // view层


## JWT

    req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzk1MTEzMjIsImlhdCI6MTU3OTUxMDEyMiwiaXNzIjoiaXJpcyIsInVzZXJJZCI6MSwidXNlck5hbWUiOiIxMTEifQ.ifxdsBxc7cyekPXnmMYLc1P2Wur7ssGiI51Q45WPass")