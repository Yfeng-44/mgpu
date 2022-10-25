#!/bin/bash
cd samples
cd spmv
./spmv -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -dim=2097152 -sparsity=0.00001 