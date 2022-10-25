#!/bin/bash
cd samples
cd shoc-reduction
./shoc-reduction -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -Size=67108864 -Iterations=2 