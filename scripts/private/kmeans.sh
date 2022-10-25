#!/bin/bash
cd samples
cd kmeans
./kmeans -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -points=524288 -features=32 -clusters=20 -max-iter=1 