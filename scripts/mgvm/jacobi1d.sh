#!/bin/bash
cd samples
cd jacobi1d
./jacobi1d -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 16384 -mem-allocator-type hslaware-32 -n=67108864 -steps=1