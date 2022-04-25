#!/bin/bash
cntNum1=0
cntNum2=0
for plan_id in `cat '/home/dev/lipeiyong/sqlRes/planIds.txt'`
do
  echo "start planId: $plan_id}"
  sub_id=`expr $plan_id % 128`
  cmd="SELECT plan_id ,count(id) as 总数, COUNT(DISTINCT(user_id)) as 发言人数 FROM xes_leave_message_log_${sub_id} WHERE 1 AND plan_id = ${plan_id} ;"
  cnt=$(mysql -ubarrage_rw -hsjhl-xy-livebarrage-mysql-rw -P5306  -pTajIEncXwarqOUdSGfuSreEfGZCfiK9+ -e "${cmd}" barrage >>  '/home/dev/lipeiyong/sqlRes/messageV2.log' )
  num3=`echo ${cnt} | awk '{print $2}'`
done
echo "finish"
exit