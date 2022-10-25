#!/bin/bash
cd samples
cd syr2k
./syr2k -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 30000000 -ni=1024 -nj=1024 