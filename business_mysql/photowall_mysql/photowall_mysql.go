package main

import (
	"fmt"
	"strconv"
	"time"
)

const(
	photoWallInfo = "xes_photo_wall_info"
	stuWallInfo = "xes_stu_wall_info"
	pwiSQL = "ALTER TABLE xes_stu_wall_info_%v MODIFY COLUMN `ai_img_result` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'ai图片识别结果 1：通过；2:不通过 ；3:ai错误 ；4:ai功能关闭' AFTER `img_url`, MODIFY COLUMN `ai_face_result` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'ai图片识别结果 1：通过；2:不通过 ；3:ai错误 ；4:ai功能关闭' AFTER `ai_img_result`,  ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)' AFTER `brand`, ADD INDEX `index_plan_id`(`plan_id`) USING BTREE;"
)
func main(){
	add2()
}


func add2(){
	fmt.Println(fmt.Sprintf("ALTER TABLE xes_stu_wall_info ADD COLUMN `source` int NOT NULL DEFAULT 0 COMMENT '来源,0=课下,3=学习报告重提交';"))
	for i := 0 ; i < 12 ; i++{
		t := time.Unix(1625252583 , 0)
		t = t.AddDate(0 , i , 0)

		fmt.Println(fmt.Sprintf("ALTER TABLE xes_stu_wall_info_%v ADD COLUMN `source` int NOT NULL DEFAULT '' COMMENT '来源,0=课下,3=学习报告重提交';",GetTableSuffixByMonth(t.Unix())))
	}
}

func add1(){
	fmt.Println(fmt.Sprintf("ALTER TABLE xes_photo_wall_info ADD COLUMN `business_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '业务类型,1=拍照作答，2=一起实验' AFTER `businessline_id`;"))
	fmt.Println(fmt.Sprintf("ALTER TABLE xes_stu_wall_info ADD COLUMN `business_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '业务类型,1=拍照作答，2=一起实验' AFTER `businessline_id`;"))

	for i := 0 ; i < 12 ; i++{
		t := time.Unix(1625252583 , 0)
		t = t.AddDate(0 , i , 0)

		fmt.Println(fmt.Sprintf("ALTER TABLE xes_stu_wall_info_%v ADD COLUMN `class_name` varchar(255) NOT NULL DEFAULT '' COMMENT '班级名称';",GetTableSuffixByMonth(t.Unix())))
	}
	//for i := 0 ; i < 6 ; i++{
	//	t := time.Unix(1633201383 , 0)
	//	t = t.AddDate(0 , i , 0)
	//	fmt.Println(fmt.Sprintf("ALTER TABLE xes_photo_wall_info_%v ADD COLUMN `business_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '业务类型,1=拍照作答，2=一起实验' AFTER `businessline_id`;",GetTableSuffixByMonth(t.Unix())))
	//}
}


func GetTableSuffixByMonth(sTime int64) string {
	t := time.Unix(sTime, 0)
	year := strconv.Itoa(t.Year())
	month := t.Format("01")
	return year + month
}