#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table transaction (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    fromwhom varchar(50),
    towhom varchar(50),
    message varchar(200),
    point int(10)
    );"

$CMD_MYSQL -e "create table user (
id int(10)  AUTO_INCREMENT NOT NULL primary key,
name varchar(50),
);"

# $CMD_MYSQL -e  "insert into transaction values (, '記事1', '記事1です。');"
# $CMD_MYSQL -e  "insert into article values (2, '記事2', '記事2です。');"