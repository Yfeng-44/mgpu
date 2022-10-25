#!/bin/bash
cd samples
cd kmeans
./kmeans -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -points=524288 -features=32 -clusters=20 -max-iter=1 