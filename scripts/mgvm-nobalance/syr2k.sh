#!/bin/bash
cd samples
cd syr2k
./syr2k -timing -no-progress-bar -report-all -scheduling lasp -platform-type customtlb -use-lasp-hsl-mem-alloc -sched-partition Xdiv -custom-hsl 512 -mem-allocator-type hslaware-1 -max-inst 30000000 -ni=1024 -nj=1024 