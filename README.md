# go-gin-learn

## 本地环境启动
### 数据库
```
docker pull mysql 
mkdir -p /data/docker-mysql/  
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<self define> -v /data/docker-mysql:/var/lib/mysql -d mysql  
创建数据库 导入sql
```



## docker应用部署
### 镜像
```
修改dockerfile的host为mysql:3306  
docker build -t gin-blog-docker .
```
### 容器
```
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker
```