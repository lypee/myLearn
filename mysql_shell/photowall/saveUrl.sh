#!/bin/bash
cntNum1=0
cntNum2=0
for plan_id in `cat '/home/dev/lipeiyong/photowall-mysql/planIds.txt'`
do
  echo "start planId: $plan_id}"
  sub_id=`expr $plan_id % 128`
  cmd="SELECT  img_url ,ai_img_result,ai_face_result,ai_img_code,ai_face_code,ai_lightness FROM xes_stu_wall_info_202111 WHERE 1 AND plan_id = ${plan_id} ;"
  cnt=$(mysql -uphotowall_rw -hsjhl-xy-livesrv-mysql-rw -P4306  -pnBWxhG947MIobHaMtzGrTPNRzzGIDBa_ -e "${cmd}" live_photo_wall >>  '/home/dev/lipeiyong/photowall-mysql/message.log' )
  num3=`echo ${cnt} | awk '{print $2}'`
done
echo "finish"
exit

