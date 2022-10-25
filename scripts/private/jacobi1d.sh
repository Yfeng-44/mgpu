#!/bin/bash
cd samples
cd jacobi1d
./jacobi1d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -n=67108864 -steps=1