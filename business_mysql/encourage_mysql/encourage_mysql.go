package main

import (
	"fmt"
)

const(
	photoWallInfo = "xes_photo_wall_info"
	stuWallInfo = "xes_stu_wall_info"
	pwiSQL = "ALTER TABLE xes_stu_wall_info_%v MODIFY COLUMN `ai_img_result` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'ai图片识别结果 1：通过；2:不通过 ；3:ai错误 ；4:ai功能关闭' AFTER `img_url`, MODIFY COLUMN `ai_face_result` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'ai图片识别结果 1：通过；2:不通过 ；3:ai错误 ；4:ai功能关闭' AFTER `ai_img_result`,  ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)' AFTER `brand`, ADD INDEX `index_plan_id`(`plan_id`) USING BTREE;"
)
func main(){
	add3()
}

// 卡牌背包新增
func add3(){
	base := "ALTER TABLE `live_encourage`.`xes_encourage_backpack_%d`\nADD COLUMN `last_gain_time` int(11) NOT NULL COMMENT '最新获得时间' AFTER `stuCouId`;"
	for i := 0 ; i< 64 ; i++{
		fmt.Println(fmt.Sprintf(base , i))
	}
}

