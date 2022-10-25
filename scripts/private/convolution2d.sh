#!/bin/bash
cd samples
cd convolution2d
./convolution2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Ydiv -ni=8192 -nj=8192 