#!/bin/bash
cd samples
cd pagerank
./pagerank -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 8192 -mem-allocator-type hslaware-16 -node=8192 -sparsity=0.5 -iterations=1 