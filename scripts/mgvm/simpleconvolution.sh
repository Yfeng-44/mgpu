#!/bin/bash
cd samples
cd simpleconvolution
./simpleconvolution -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 16384 -mem-allocator-type hslaware-32 -width=8190 -height=8190 