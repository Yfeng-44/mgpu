#!/bin/bash
cd samples
cd spmv
./spmv -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 512 -mem-allocator-type hslaware-1 -dim=2097152 -sparsity=0.00001 