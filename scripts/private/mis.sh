#!/bin/bash
cd samples
cd mis
./mis -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -numNodes=524288 -numItems=1048576 