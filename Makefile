NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
ECHO=echo
else
ECHO=/bin/echo -e
endif

all: 
	@$(ECHO) "$(OK_COLOR)==> Building$(NO_COLOR)"
	go get -v ./...
	go test -v

bin: 
	@$(ECHO) "$(OK_COLOR)==> Building$(NO_COLOR)"
	go build

clean:
	@rm -rf dist/ packer-post-processor-vagrant-vmware-ovf

format:
	go fmt ./...

dist:
	@$(ECHO) "$(OK_COLOR)==> Building Packages...$(NO_COLOR)"
	@gox -osarch="darwin/386 darwin/amd64 linux/386 linux/amd64 freebsd/386 freebsd/amd64 openbsd/386 openbsd/amd64 windows/386 windows/amd64 netbsd/386 netbsd/amd64"
	@mv packer-post-processor-vagrant-vmware-ovf_darwin_386 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.darwin-i386.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_darwin_amd64 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.darwin-amd64.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_freebsd_386 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.freebsd-i386.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_freebsd_amd64 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.freebsd-amd64.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_linux_386 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.linux-i386.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_linux_amd64 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.linux-amd64.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_netbsd_386 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.netbsd-i386.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_netbsd_amd64 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.netbsd-amd64.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_openbsd_386 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.openbsd-i386.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_openbsd_amd64 packer-post-processor-vagrant-vmware-ovf; tar cvfz packer-post-processor-vagrant-vmware-ovf.openbsd-amd64.tar.gz packer-post-processor-vagrant-vmware-ovf; rm packer-post-processor-vagrant-vmware-ovf
	@mv packer-post-processor-vagrant-vmware-ovf_windows_386.exe packer-post-processor-vagrant-vmware-ovf.exe; zip packer-post-processor-vagrant-vmware-ovf.windows-i386.zip packer-post-processor-vagrant-vmware-ovf.exe; rm packer-post-processor-vagrant-vmware-ovf.exe
	@mv packer-post-processor-vagrant-vmware-ovf_windows_amd64.exe packer-post-processor-vagrant-vmware-ovf.exe; zip packer-post-processor-vagrant-vmware-ovf.windows-amd64.zip packer-post-processor-vagrant-vmware-ovf.exe; rm packer-post-processor-vagrant-vmware-ovf.exe
	@mkdir -p dist/
	@mv packer-post-processor-vagrant-vmware-ovf* dist/.

.PHONY: all clean deps format test updatedeps
