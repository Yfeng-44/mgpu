all: esc.go

esc.go: kernels.hsaco
	esc -o $@ -pkg correlation -private $^

kernels.disasm: kernels.hsaco
	/opt/rocm/hcc/bin/llvm-objdump -mcpu=fiji -disassemble $^ > $@

kernels.hsaco: native/correlation.cl
	clang-ocl -mcpu=fiji -o $@ $^
