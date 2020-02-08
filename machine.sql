create database if not exists machine
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;
USE machine;
SET NAMES utf8;

drop table if exists machine_info;
CREATE TABLE `machine_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '机器的名字',
  `resume` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '机器的简介',
  `type` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '机器类型',
  `site` int(11) NOT NULL DEFAULT '0' COMMENT '机位',
  `available` BOOLEAN NOT NULL COMMENT '设备是否可用',
  `bind` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '绑定的标识，比如AGENT ID',
  `ctime` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='机器信息';

drop table if exists order_info;
CREATE TABLE `order_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `mid` int(11) NOT NULL DEFAULT '0' COMMENT '对应的machine id',
  `po`  VARCHAR(30) NOT NULL DEFAULT '',
  `mo`  VARCHAR(30) NOT NULL DEFAULT '',
  `lot`  VARCHAR(30) NOT NULL DEFAULT '',
  `status` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '该订单是否cancel、done、going、wait，只能有一个going',
  `ctime` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `ix_mid` (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单的信息';

drop table if exists class_info;
CREATE TABLE `class_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `mid` int(11) NOT NULL DEFAULT '0' COMMENT '对应的machine id',
  `name` VARCHAR(25) NOT NULL DEFAULT '' COMMENT '节点的名字',
  `resume` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '机器的简介',
  `number` int(11) NOT NULL DEFAULT '0' COMMENT '节点编号',
  `available` BOOLEAN NOT NULL COMMENT '设备是否可用',
  `ctime` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`),
	KEY `ix_mid` (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单的信息';
