# What is this
M1 Mac でDocker-Mysql環境を作成

# bootstrap
```sh
docker pull --platform linux/x86_64 mysql:latest
```

### DLしていなければ以下DLが走る
```sh
Unable to find image 'mysql:latest' locally
latest: Pulling from library/mysql
a4b007099961: Pull complete 
e2b610d88fd9: Pull complete 
38567843b438: Pull complete 
5fc423bf9558: Pull complete 
aa8241dfe828: Pull complete 
cc662311610e: Pull complete 
9832d1192cf2: Pull complete 
...
.......
```

### imageを確認
```sh
docker images
```

latestタグで確認できる
```sh
REPOSITORY        TAG       IMAGE ID       CREATED        SIZE
mysql             latest    562c9bc24a08   5 days ago     521MB
```

##＃ 起動
```sh
docker run --name mysql-test -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql
```

# 接続

## docker 接続
```sh
docker exec -it mysql-test bash
```

## mysql 接続
```sh
root@1594af897019:/# mysql -uroot -p
Enter password: 123456
```

```sh
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.28 MySQL Community Server - GPL

Copyright (c) 2000, 2022, Oracle and/or its affiliates.
```
