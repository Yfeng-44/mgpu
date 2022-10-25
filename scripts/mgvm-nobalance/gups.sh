#!/bin/bash
cd samples
cd gups
./gups -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -sched-partition Xdiv -custom-hsl 1024 -mem-allocator-type hslaware-2  