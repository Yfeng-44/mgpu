#!/bin/bash
cd samples
cd matrixtranspose
./matrixtranspose -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -width=2048 