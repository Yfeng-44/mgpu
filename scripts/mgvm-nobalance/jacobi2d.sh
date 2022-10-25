#!/bin/bash
cd samples
cd jacobi2d
./jacobi2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -sched-partition Ydiv -custom-hsl 4096 -mem-allocator-type hslaware-8 -n=4096 -steps=1