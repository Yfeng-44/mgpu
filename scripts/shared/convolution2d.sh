#!/bin/bash
cd samples
cd convolution2d
./convolution2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Ydiv -ni=8192 -nj=8192 