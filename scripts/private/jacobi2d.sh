#!/bin/bash
cd samples
cd jacobi2d
./jacobi2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Ydiv -n=4096 -steps=1