#!/bin/bash
cd samples
cd fastwalshtransform
./fastwalshtransform -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -length=8388608 