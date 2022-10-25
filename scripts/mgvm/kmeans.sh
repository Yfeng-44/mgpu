#!/bin/bash
cd samples
cd kmeans
./kmeans -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 4096 -mem-allocator-type hslaware-8 -points=524288 -features=32 -clusters=20 -max-iter=1 