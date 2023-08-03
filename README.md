

> http://www.baker-yuan.cn/articles/395

# Go 整洁模板

> https://github.com/evrone/go-clean-template

- [go-clean-template](./go-clean-template)



# project-layout

> https://github.com/golang-standards/project-layout

- [project-layout](./project-layout)

# kratos

> https://github.com/go-kratos/kratos-layout
>
> https://github.com/go-kratos/examples/blob/main/README-CN.md
>
> https://github.com/go-kratos/examples

```bash
kratos new kratos-layout

$ cd kratos-layout
$ go generate ./...
$ go build -o ./bin/ ./... 
$ ./bin/kratos-layout -conf ./configs

curl 'http://127.0.0.1:8000/helloworld/kratos'
```



- [kratos-layout](./kratos-layout)
- [kratos-layout-blog](./kratos-layout-blog)



# nunu

- [nunu-layout-basic](./nunu-layout-basic)
- [nunu-layout-advanced](./nunu-layout-advanced)

> https://github.com/go-nunu/nunu/blob/main/README_zh.md

```bash
nunu new nunu-layout-advanced

› cd nunu-layout-advanced 
› nunu run 

Basic Layout 包含一个非常精简的架构目录结构，适合非常熟悉Nunu项目的开发者使用。
```



```bash
nunu new nunu-layout-basic

› cd nunu-layout-basic 
› nunu run 

Advanced Layout 包含了很多Nunu的用法示例（ db、redis、 jwt、 cron、 migration等），适合开发者快速学习了解Nunu的架构思想。此命令将创建一个名为projectName的目录，并在其中生成一个优雅的Golang项目结构。
```



# go-zero

> https://go-zero.dev/docs/concepts/layout
>
> https://github.com/zeromicro/zero-examples



# 其他

> https://zhuanlan.zhihu.com/p/399095776