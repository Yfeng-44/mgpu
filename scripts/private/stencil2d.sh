#!/bin/bash
cd samples
cd stencil2d
./stencil2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -row=2048 -col=2048 