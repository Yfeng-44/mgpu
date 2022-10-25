#!/bin/bash
cd samples
cd syrk
./syrk -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -max-inst 10000000 -ni=2048 -nj=2048 