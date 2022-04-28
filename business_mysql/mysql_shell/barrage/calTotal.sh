#!/bin/bash
cntNum1=0
cntNum2=0
for plan_id in `cat '/home/dev/lipeiyong/sqlRes/planIds.txt'`
do
  echo "start planId: $plan_id}"
  sub_id=`expr $plan_id % 128`
  cmd="SELECT count(id), COUNT(DISTINCT(user_id))  FROM xes_leave_message_log_${sub_id} WHERE 1 AND plan_id = ${plan_id} AND create_time BETWEEN '2021-10-04 00:01:01'  and '2021-10-07 23:59:59' ;"
  cnt=$(mysql -ubarrage_rw -hsjhl-xy-livebarrage-mysql-rw -P5306  -pTajIEncXwarqOUdSGfuSreEfGZCfiK9+   -e "${cmd}" barrage)
  num3=`echo ${cnt} | awk '{print $3}'`
  num4=`echo ${cnt} | awk '{print $4}'`
  cntNum1=$((${cntNum1} + $num3))
  cntNum2=$((${cntNum2} + $num4))
done



echo "1 $cntNum1 "
echo "2 $cntNum2"

exit
