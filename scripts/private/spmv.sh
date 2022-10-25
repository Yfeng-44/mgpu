#!/bin/bash
cd samples
cd spmv
./spmv -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -dim=2097152 -sparsity=0.00001 