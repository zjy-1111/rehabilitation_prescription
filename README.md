# rehabilitation_prescription
康复处方
#### 数据库设计
1. 预约表
```
CREATE TABLE `reservation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `time` int(10) unsigned NOT NULL COMMENT ‘预约时间’,
  `doctor_name` varchar(100) DEFAULT '' COMMENT ‘预约医生’,
  `address` varchar(100) DEFAULT '' COMMENT ‘预约地点’,
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='预约订单';

```
