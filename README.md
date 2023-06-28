# GraduateProgramMSI
毕业课题的分析代码和网站代码

## 运行环境

Windows

## Analysis文件夹

* 该文件夹中是论文的分析和绘图代码，需要在jupyter notebook中查看
* 依赖包：python3、sklearn==1.2.1、pandas、numpy、matplotlib、seaborn、statannotations、dtreeviz、imblearn、eli5

## Website文件夹

* Website是论文中搭建的网站平台，IDE使用vscode开发
* 开发语言：前端javascript、后端golang、检测python3、数据库mysql
* python检测脚本依赖包：joblib、pandas、sklearn==1.2.1、eli5
* 后端开发依赖见go.mod文件
* config.ini文件配置项目端口和数据库端口、账户、密码
* 使用`go run /path/main.go`运行脚本
* 使用`go mod download`安装依赖
* 检测需要使用的基因集['MLH1', 'EPM2AIP1','MSH4','RPL22L1', 'HSPH1','TRIM7', 'PRTFDC1','SMAP1', 'DDX27','CXCL9', 'WDTC1','FECH',  'SFXN1','COL18A1']
