#!/bin/bash
cd samples
cd shoc-reduction
./shoc-reduction -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -Size=67108864 -Iterations=2 