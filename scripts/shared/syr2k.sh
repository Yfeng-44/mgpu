#!/bin/bash
cd samples
cd syr2k
./syr2k -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -max-inst 30000000 -ni=1024 -nj=1024 