#!/bin/bash
cd samples
cd syrk
./syrk -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 1024 -mem-allocator-type hslaware-2 -max-inst 10000000 -ni=2048 -nj=2048 