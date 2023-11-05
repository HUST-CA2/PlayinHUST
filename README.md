# PlayinHUST 玩在华科
"玩在华科"网站项目

### 如何协作？
请参考这篇博客，向项目提交Pull Request  
[github中Pull Request的使用](https://tsailooy.top/2023/10/30/github%E4%B8%ADPullRequest%E7%9A%84%E4%BD%BF%E7%94%A8/)

### 后端的同志如何运行代码
使用`git clone`到本地后，无法直接运行，需要添加配置文件`conf.go`，设置数据库配置和服务器配置

`conf.go`内容如下，创建后放在`/common`文件夹下
```go
package common

import "github.com/jinzhu/gorm"

//这个文件用来配置

//服务器配置
var (
	ServerName string = "localhost"
	Port       string = ":8080"
)

//数据库配置
var (
	dbdriver      = "mysql"
	dbusername    = "username"
	dbpassword    = "password"
	dbaddr        = "(localhost:3306)"
	dbname        = "playinhustdb"
	dboption      = "multiStatements=true&&parseTime=true"
	PlayinHUSTDB  *gorm.DB
	sqlConnection = dbusername + ":" + dbpassword + "@tcp" + dbaddr + "/" + dbname + "?" + dboption
)
```

然后添加`./gitignore`文件
```./gitignore
# common下的conf.go为配置文件

./common/conf.go
```

设置好后，打开终端，输入`go run main.go`即可在本地运行

### 前端的同志在哪修改代码

`./template`下存放各个页面模版html，`./public`下存放所需静态资源img,js和css

若需要在html中引用资源，则后端已将相对地址`./public/resource/`映射为`/src/`，以logo为例，引用的相对地址为`/src/img/CAlogo.png`

```html
<div class="logo"><a href="index.html"><img src="/src/img/CAlogo.png" alt="Logo"></a></div>
```

### 具体细节可以直接在飞书沟通