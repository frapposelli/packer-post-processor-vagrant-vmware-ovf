# packer-post-processor-vagrant-vmware-ovf [![Build Status](https://travis-ci.org/gosddc/packer-post-processor-vagrant-vmware-ovf.svg)](https://travis-ci.org/gosddc/packer-post-processor-vagrant-vmware-ovf)

This packer plugin leverages [VMware OVF Tool](http://www.vmware.com/support/developer/ovf) to create a ```vmware_ovf``` Vagrant box that is compatible with [vagrant-vcloud](https://github.com/frapposelli/vagrant-vcloud), [vagrant-vcenter](https://github.com/gosddc/vagrant-vcenter) and [vagrant-vcloudair](https://github.com/gosddc/vagrant-vcloudair), you can find a detailed explanation of the format [here on the wiki](https://github.com/gosddc/packer-post-processor-vagrant-vmware-ovf/wiki/vmware_ovf-Box-Format).

This plugin is used to build the boxes available at https://vagrantcloud.com/gosddc you can find the [packer templates here](https://github.com/gosddc/packer-templates)

## Prerequisites

Software:

  * VMware OVF Tool

Notes:

  * This post processor only works with the VMware builder.

## Installation

Starting from Packer v0.7.0 there are new ways of installing plugins, [see the official Packer documentation](http://www.packer.io/docs/extend/plugins.html) for further instructions.

## Usage

In your JSON template add the following post processor:

```json
  "post-processors": [
    {
        "type": "vagrant-vmware-ovf"
    }
  ]
```

Other parameters available are:

- ```provider```: You can override the provider metadata to ```vcloud``` or ```vcenter``` to build legacy boxes for old [vagrant-vcloud](https://github.com/frapposelli/vagrant-vcloud) and [vagrant-vcenter](https://github.com/gosddc/vagrant-vcenter) installs.
- ```compression```: You can set compression of the box with an integer from 0 to 9 (default is 6).

If you don't want to compile the code, you can [grab a release here](https://github.com/gosddc/packer-post-processor-vagrant-vmware-ovf/releases).
