#!/bin/bash
cd samples
cd mis
./mis -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 512 -mem-allocator-type hslaware-1 -numNodes=524288 -numItems=1048576 