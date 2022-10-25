#!/bin/bash
cd samples
cd convolution2d
./convolution2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Ydiv -custom-hsl 16384 -mem-allocator-type hslaware-32 -ni=8192 -nj=8192 