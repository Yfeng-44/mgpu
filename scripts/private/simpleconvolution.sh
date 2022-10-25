#!/bin/bash
cd samples
cd simpleconvolution
./simpleconvolution -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -width=8190 -height=8190 