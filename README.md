# GitHub Imgur

自制图床网站, 使用 GitHub 作为图床

[项目演示地址](https://imgur.chenxutalk.top)

# 技术选型

前端: html + css + js

后端: go + gin

# 使用

拉取代码

```sh
git clone https://github.com/Guo-Chenxu/github-imgur.git
```

### 前端

[前端](./imgur-front/)

#### 本地运行

直接打开 `index.html`

#### 部署到服务器

在 `nginx` 的配置文件中将请求指向 `index.html`

### 后端

[后端](./imgur-backend/)

自行填写 `conf/config.ini` 中的配置

#### 本地运行

```sh
go mod tidy
go run main.go
```

#### 部署到服务器

```sh
go build
```

将生成的可执行文件 `imgur-backend`, 以及 `conf/config.ini` _(或者可以在运行时指定配置文件)_ 上传到服务器

```sh
./imgur-backend
```

在 `nignx` 配置文件中设置路径转发

# 实现功能

-   图片上传到 github, 并返回图床链接
-   自动将返回链接写入到剪切板中
-   适配 PC 端和移动端

# 效果展示

<img src="https://cdn.jsdelivr.net/gh/Guo-Chenxu/imgs@main/imgs/202309201613786.png"/>

<img src="https://cdn.jsdelivr.net/gh/Guo-Chenxu/imgs@main/imgs/202309201614029.png"/>

# 参考

前端参考自: [img_compress_rotate_preview_upload](https://github.com/legend-li/img_compress_rotate_preview_upload) _(删除了前端的压缩操作)_
