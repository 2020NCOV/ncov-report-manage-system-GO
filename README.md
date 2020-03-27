# ncov-report-manage-system-GO

> 疫情防控--后台管理系统-Go后端
>
> 现在只是一个简单的小demo，大家一起完善

## 项目导航

- ncov-report-manage-system-GO
  - [项目导航](#项目导航)
  - [框架](#框架)
  - [目录结构](#目录结构)
  - [本地配置](#本地配置)
  - [子目录文档](#子目录文档)
    - [pages](#pages)
    - [model](#model)
  - [需求](#需求)

## 框架

- gin框架([文档](https://github.com/gin-gonic/gin))
- go-admin后台管理框架([地址](https://github.com/GoAdminGroup/go-admin))

## 目录结构

```
|-- ncov-report-manage-system-GO
    |-- config							//各种配置文件在此文件夹中
    |   |-- config.go
    |-- controller					    //此文件夹中写逻辑代码，也可自建符合逻辑的文件夹
    |-- model							//详见model文件夹中的README.md文档
    |   |-- README.md				
    |   |-- admin_user.go
    |   |-- org_dep.go
    |   |-- org_whitelist.go
    |   |-- organization.go
    |   |-- report_record_default.go
    |   |-- tables.go
    |   |-- wx_mp_bind_info.go
    |-- pages							//详见pages文件夹中的README.md文档
    |   |-- DashBoard.go
    |   |-- README.md				
    |-- server							//此文件夹功能参照上个项目中的server和api文件夹
    |   |-- apiFunc.go
    |   |-- main.go
    |-- main.go							//程序入口
    |-- database						//数据库文件
        |-- go-admin.sql				//go-admin框架需要
        |-- ncov_db_v1.0.0.sql			//项目需要（主仓库中的数据库文件）
```

 ### 本地配置

- **fork本仓库，并clone到本地**

- **Goland中导入该项目**

- **修改数据库配置信息**

新建文件夹config.json，仿照config.example.json文件进行配置

```
{
  "database": {
    "default": {
      "host": "127.0.0.1",
      "port": "3306",
      "user": "数据库用户名",
      "pwd": "用户密码",
      "name": "数据库名",
      "max_idle_con": 50,
      "max_open_con": 150,
      "driver": "mysql"
    }
  },
  "prefix": "ncov",  //注意此处改为ncov,为网站前缀：http://127.0.0.1:9033/ncov/...
  "store": {
    "path": "./uploads",
    "prefix": "uploads"
  },
  "language": "en"
  "title": "ncov",
  "login_title": "疫情上报系统",
  "logo": "<p2><b>N</b>COV<p2>",
  "mini_logo": "<p2><b>N</b>C<p2>"
}
```


- **将database文件夹中的.sql文件导入数据库中**

注：两个都需要导入，如果之前已经导入过张老师提供的那个数据库，就只需将go-admin框架需要的数据库go-admin.sql导入即可

- **运行项目**

```
go run main.go

或者右键main.go运行
（mac为两只手指点击）
```

- **浏览器访问**

```
http://127.0.0.1:9033/ncov/login
```
用户名：admin（总管理员）
密码：admin

用户名：operator(管理员）
密码：admin

- 然后就开始开发吧！

## 子目录文档

### pages

- 此文件夹中只有一个文件DashBoard.go，主页的页面展示（相当于前端页面）
- 对主页的统计信息进行修改主要在此文件中，其他具体逻辑代码请在controller文件夹中新建文件
- 或者自行新建文件夹（需符合框架整体逻辑）

### model

>
>
>此文件夹中的文件使用插件adm从数据库中的表中自动生成的文件，详见[文档](https://gitee.com/go-admin/docs/wikis/pages?sort_id=1888257&doc_id=516564)

#### 1. 访问url见model/table.go文件

     http://127.0.0.1:9033/ncov/info/xxx(表名)

#### 2. 自动生成的是缺省表格，需要修改model/xxx.go这些文件来编辑对应的表格的样式

   在model/wx_mp_bind_info中有做一些简单的修改，model中其余文件都是自动生成的原始文件

​	如何编辑可参考[官方框架小demo](https://github.com/GoAdminGroup/go-admin/tree/master/examples/datamodel)的相关文件以及[demo网页展示](https://demo.go-admin.com/admin)

## 需求

>请移步issue认领
>
>认领方式：在issue下面评论
>
>提交PR：如果你的PR已经解决这个issue，请在提交PR同时备注对应issue，以便merge之后自动关闭issue（可使用#加对应issue的id，比如：#43）

#### 部门管理员操作的设置

- 使不同用户登录进去显示不同内容
- 不同用户操作范围不同

#### 白名单、小程序绑定用户、未提交名单的匹配问题

- 未提交名单的自动更新
- 匹配

#### 数据下载与数据上传

- 数据上传主要是白名单导入、用户头像上传
- 数据下载主要是每日上报数据的下载、用户表的下载等等
- 框架已提供下载功能，注意不同用户的下载权限

#### 提醒功能

- 有紧急情况时（比如有人体温超过多少时），提醒管理员
- 提醒方式自行调研（短信、微信、邮件等等）

#### 主页信息统计（图表形式显示）（可多人认领）

1. **上报信息统计**
2. **疫情概况统计**
3. **体温情况统计**
4. **人员分布统计**
5. **返回本地人员统计**
6. **心理健康信息统计**

> 统一要求：
>
> - 使用图表形式直观地显示信息
> - 点击图表或其他形式还可看到详情页，展示详细信息（自行发挥）

#### 登录验证

- 图形验证、滑动滑块验证、验证码验证等等

#### 图表处理（可多人认领）

- 需要处理使用插件生成的默认表格，将表格样式统一成我们自己的风格，需要处理model中所有表格
- 主要处理：图标titile、图标简介description、列的增删、表头需要改为中文的也需要修改、添加功能键等等（其他需求自行发挥）

1. org_whitelist
2. admin_user
3. org_dep
4. organization
5. wx_mp_bind_info
6. report_record_default
