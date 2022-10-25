#!/bin/bash
cd samples
cd fastwalshtransform
./fastwalshtransform -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -switch-tlb-striping -sched-partition Xdiv -custom-hsl 2048 -mem-allocator-type hslaware-4 -length=8388608 