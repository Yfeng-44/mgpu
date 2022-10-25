#!/bin/bash
cd samples
cd pagerank
./pagerank -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -node=8192 -sparsity=0.5 -iterations=1 