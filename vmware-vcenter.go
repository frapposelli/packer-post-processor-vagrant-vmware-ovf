package main

import (
	"fmt"
	"github.com/mitchellh/packer/packer"
	"path/filepath"
  "os"
  "os/exec"
  "io/ioutil"
  "strings"
)

type VMwarevCenterProvider struct{}

func (p *VMwarevCenterProvider) KeepInputArtifact() bool {
	return false
}

func (p *VMwarevCenterProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	// Create the metadata
	metadata = map[string]interface{}{"provider": "vcenter"}

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
    
  // start upload
  ui.Message(fmt.Sprintf("ovftool is going create an ovf out of %s", vmx))
  
  program := "ovftool"
  sourcetype := "--sourceType=VMX"
  targettype := "--targetType=OVF"

  // start upload
  ui.Message(fmt.Sprintf("this is our ovftool run %s %s %s %s %s", program, sourcetype, targettype, vmx, basepath + "/" + ovf))

  ui.Message(fmt.Sprintf("Creating directory: %s", basepath))

  if err := os.Mkdir(basepath, 0755); err != nil { 
    ui.Message(fmt.Sprintf("err: %s", err))
  } 

  cmd := exec.Command(program, sourcetype, targettype, vmx, basepath + "/" + ovf)

	ui.Message(fmt.Sprintf("Starting ovftool"))

	cmd.Start()
	cmd.Wait()

	ui.Message(fmt.Sprintf("Reading files in %s", basepath))
  files, _ := ioutil.ReadDir(basepath)
  for _, path := range files {
  //         ui.Message(fmt.Sprintf("%s", f.Name()))
  // }


	// Copy all of the original contents into the temporary directory
	// for _, path := range artifact.Files() {
		ui.Message(fmt.Sprintf("Copying: %s", path.Name()))

		dstPath := filepath.Join(dir, path.Name())
		if err = CopyContents(dstPath, basepath + "/" + path.Name()); err != nil {
			return
		}
	}

	return
}
