# packer-post-processor-vagrant-vmware-ovf

This packer plugin leverages [VMware OVF Tool](http://www.vmware.com/support/developer/ovf) to create a [vagrant-vcloud](https://github.com/frapposelli/vagrant-vcloud) and [vagrant-vcenter](https://github.com/gosddc/vagrant-vcenter) compatible boxes.

This plugin is used to build the boxes available at http://vagrant.gosddc.com/boxes you can find the packer templates at https://github.com/gosddc/packer-templates

## Prerequisites

Software:

  * VMware OVF Tool
  
Notes:

  * This post processor only works with VMware

## Installation

Add

```
{
  "post-processors": {
    "vagrant-vmware-ovf": "packer-post-processor-vagrant-vmware-ovf"
  }
}
```

to your packer configuration (see: http://www.packer.io/docs/other/core-configuration.html -> Core Configuration)

Make sure that the directory which contains the packer-post-processor-vagrant-vmware-ovf executable is your PATH environmental variable (see http://www.packer.io/docs/extend/plugins.html -> Installing Plugins)

## Usage

In your JSON template add the following post processor:

```
  "post-processors": [
    {
        "type": "vagrant-vmware-ovf",
        "provider": "vcenter"
    }
  ]
```

You can change ```provider``` to ```vcloud``` to make it compatible with [vagrant-vcloud](https://github.com/frapposelli/vagrant-vcloud).