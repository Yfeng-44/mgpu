#!/bin/bash
cd samples
cd mis
./mis -timing -no-progress-bar -report-all -scheduling lasp -platform-type xortlb -mem-allocator-type lasp -use-lasp-mem-alloc -l2-tlb-striping 1 -sched-partition Xdiv -numNodes=524288 -numItems=1048576 