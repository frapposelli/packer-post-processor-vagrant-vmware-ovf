package main

import (
	"fmt"
	vmwcommon "github.com/mitchellh/packer/builder/vmware/common"
	"github.com/mitchellh/packer/packer"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type VMwareOVFProvider struct{}

func (p *VMwareOVFProvider) KeepInputArtifact() bool {
	return false
}

func (p *VMwareOVFProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	// Create the metadata
	metadata = map[string]interface{}{"provider": "vmware_ovf"}

	vmx := ""
	ovf := ""
	basepath := ""
	for _, path := range artifact.Files() {
		if strings.HasSuffix(path, ".vmx") {
			vmx = path
			ovf = filepath.Base(vmx[:len(vmx)-4] + ".ovf")
			basepath = filepath.Dir(path) + "/ovf"
		}
	}

	vmxData, err := vmwcommon.ReadVMX(vmx)
	if err != nil {
		ui.Message(fmt.Sprintf("err: %s", err))
	}

	for k, _ := range vmxData {
		if strings.HasPrefix(k, "floppy0.") {
			ui.Message(fmt.Sprintf("Deleting key: %s", k))
			delete(vmxData, k)
		}
		if strings.HasPrefix(k, "ide1:0.file") {
			ui.Message(fmt.Sprintf("Deleting key: %s", k))
			delete(vmxData, k)
		}
	}

	// remove floppy (again)
	ui.Message(fmt.Sprintf("Setting key: floppy0.present = FALSE"))
	vmxData["floppy0.present"] = "FALSE"

	// detach DVD (again)
	ui.Message(fmt.Sprintf("Setting key: ide1:0.present = FALSE"))
	vmxData["ide1:0.present"] = "FALSE"

	// Rewrite the VMX
	if err := vmwcommon.WriteVMX(vmx, vmxData); err != nil {
		ui.Message(fmt.Sprintf("err: %s", err))
	}

	program, err := FindOvfTool()
	sourcetype := "--sourceType=VMX"
	targettype := "--targetType=OVF"

	ui.Message(fmt.Sprintf("Creating directory: %s", basepath))

	if err := os.Mkdir(basepath, 0755); err != nil {
		ui.Message(fmt.Sprintf("err: %s", err))
	}

	cmd := exec.Command(program, sourcetype, targettype, vmx, basepath+"/"+ovf)

	ui.Message(fmt.Sprintf("Starting ovftool"))

	cmd.Start()
	cmd.Wait()

	ui.Message(fmt.Sprintf("Reading files in %s", basepath))
	files, _ := ioutil.ReadDir(basepath)
	for _, path := range files {
		ui.Message(fmt.Sprintf("Copying: %s", path.Name()))

		dstPath := filepath.Join(dir, path.Name())
		if err = CopyContents(dstPath, basepath+"/"+path.Name()); err != nil {
			return
		}
	}

	return
}
