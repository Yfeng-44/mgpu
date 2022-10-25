#!/bin/bash
cd samples
cd shoc-reduction
./shoc-reduction -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 16384 -mem-allocator-type hslaware-32 -Size=67108864 -Iterations=2 