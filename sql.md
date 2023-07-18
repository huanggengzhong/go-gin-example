# 数据库命令集合

### 进入数据库
进入路径:cd /usr/local/mysql/bin/ 
登录命令:mysql -u root -p
mySql账号密码:
root admin123


创建数据库
blog数据库，编码为utf8_general_ci，
```go
CREATE DATABASE blog CHARACTER SET utf8 COLLATE utf8_general_ci;

```



创建标签表
```go
CREATE TABLE `blog_tag`(
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(100) DEFAULT '' COMMENT '标签名称',
`create_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`create_by` varchar(100) DEFAULT '' COMMENT '创建人',
`modified_on` int(10)unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on` int(10)unsigned unsigned DEFAULT '0',
`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0禁用 1启用',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```

文章表
```js
CREATE TABLE `blog_article`(
`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`tag_id` int(10) UNSIGNED DEFAULT '0' COMMENT '标签id',
`title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
`desc` VARCHAR(255) DEFAULT '' COMMENT '简述',
`content` text,
`create_on` int(11) DEFAULT NULL,
`create_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) UNSIGNED DEFAULT '0' COMMENT'修改时间',
`modified_by`VARCHAR(255) DEFAULT '' COMMENT'修改人',
`deleted_on` int(10) UNSIGNED DEFAULT '0',
`state` TINYINT(3) UNSIGNED DEFAULT '1'  COMMENT'状态 0禁用 1启用',
PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT '文章管理';
```
认证表
```js
CREATE TABLE `blog_auth` (
`id` int(10)UNSIGNED NOT NULL AUTO_INCREMENT,
`username` VARCHAR(50) DEFAULT '' COMMENT '账号',
`password` VARCHAR(50) DEFAULT '' COMMENT '密码',
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT '认证表';

INSERT INTO `blog`.`blog_auth` (`id`,`username`,`password`) VALUES(null,'test','test123456');
```
