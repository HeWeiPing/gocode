#!/bin/bash

#if [ $# -lt 1 ]
#	echo Usage: ./killBeeRun program_name
#then
#	exit
#fi

#pid = `ps -ef | grep $1 | grep -v grep | grep -v PPID | awk '{print $2}'`
#
#for i in $pid
#do
#	echo "kill the process [$i]"
#	kill -9 $i
#done



#kill "bee run" and "./blogProject"
pidbee=`ps -ef | grep "bee run" | grep -v grep | awk '{print $2}'`
pidpro=`ps -ef | grep "blogProject" | grep -v grep | awk '{print $2}'`
pids=`paste -d' ' <(echo "${pidbee}") <(echo "${pidpro}")`
#echo $pids

for i in ${pids}
do
	echo "kill the process [$i]"
	kill -9 $i
done
