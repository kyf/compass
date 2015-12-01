drop database if exists `compass`;
create database compass;
use compass
delimiter //

drop table if exists `admin` //
create table `admin`(
	`id` int(11) primary key auto_increment comment 'primary key',
	`admin_name` varchar(100) not null comment 'admin login name',
	`admin_pwd` varchar(250) not null comment 'admin login password'
)engine=InnoDB default charset=utf8 comment 'administrator login account'
//

drop table if exists `app` //
create table `app`(
	`id` int(11) primary key auto_increment comment 'primary key',
	`name` varchar(250) not null comment 'app name',
	`icon` varchar(250) not null default '' comment 'app icon url',
	`desc` text comment 'app description',
	`apk` varchar(300) not null default '' comment 'app apk url'
)engine=InnoDB default charset=utf8 comment 'app main table'
//

drop table if exists `setting` //
create table `setting`(
	`id` int(11) primary key auto_increment comment 'primary key',
	`ad_show` tinyint(1) not null default 0 comment 'whether show ad'
)engine=InnoDB default charset=utf8 comment 'app setting table'
//

delimiter ;
