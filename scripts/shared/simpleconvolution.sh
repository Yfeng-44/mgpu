#!/bin/bash
cd samples
cd simpleconvolution
./simpleconvolution -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -width=8190 -height=8190 