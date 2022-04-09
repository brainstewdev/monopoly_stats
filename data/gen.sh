#!/bin/bash
for ((i=1;i<=$1;i++))
do
   ../monopoly_stats -turns=45 -n=4 -json > ./jsons/game_$i.json
done