
CREATE TABLE `xes_stu_wall_info_202201` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '业务id',
  `interaction_id` varchar(100) NOT NULL DEFAULT '' COMMENT '互动id',
  `plan_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次id',
  `stu_id` int(11) NOT NULL DEFAULT '0' COMMENT '学生id',
  `stu_name` varchar(50) NOT NULL DEFAULT '' COMMENT '学生姓名',
  `stu_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '学生头像url',
  `img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片url',
  `ai_img_result` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'ai图片识别结果，0:不通过，1：通过',
  `ai_face_result` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'ai人脸识别结果 0:不通过，1:通过',
  `ai_img_code` int(11) NOT NULL DEFAULT 0 COMMENT 'ai图片返回码',
  `ai_face_code` int(11) NOT NULL DEFAULT 0 COMMENT 'ai检测人脸返回码',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `class_id` int(11) NOT NULL DEFAULT 0 COMMENT '班级ID',
  `brand` int(3) NOT NULL DEFAULT 0 COMMENT '来源',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '结束时间',
  PRIMARY KEY (`id`),
  KEY `index_interaction_id` (`interaction_id`)
) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT = '学生上传拍照上墙信息表-李佩永';


CREATE TABLE `xes_photo_wall_info_202110` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '业务id',
  `interaction_id` varchar(100) DEFAULT '0' NOT NULL COMMENT '互动id',
  `plan_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次id',
  `start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '结束时间',
  `duration` int(11) NOT NULL DEFAULT 0 COMMENT '持续时间',
  `gold` int(11) NOT NULL DEFAULT 0 COMMENT '金币数',
  `teacher_id` int(11) NOT NULL DEFAULT 0 COMMENT '发起互动的老师id',
  `teacher_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '老师类型',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `brand` int(3) NOT NULL DEFAULT '0' COMMENT '来源',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_interaction_id` (`interaction_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '拍照上墙互动表-李佩永';