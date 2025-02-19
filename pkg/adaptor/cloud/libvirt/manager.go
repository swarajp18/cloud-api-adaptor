//go:build cgo

// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package libvirt

import (
	"flag"

	"github.com/confidential-containers/cloud-api-adaptor/pkg/adaptor/cloud"
)

var libvirtcfg Config

type Manager struct{}

const (
	defaultURI            = "qemu:///system"
	defaultPoolName       = "default"
	defaultNetworkName    = "default"
	defaultDataDir        = "/var/lib/libvirt/images"
	defaultVolName        = "podvm-base.qcow2"
	defaultLaunchSecurity = ""
	defaultFirmware       = "/usr/share/edk2/ovmf/OVMF_CODE.fd"
)

func init() {
	cloud.AddCloud("libvirt", &Manager{})
}

func (*Manager) ParseCmd(flags *flag.FlagSet) {

	flags.StringVar(&libvirtcfg.URI, "uri", defaultURI, "libvirt URI")
	flags.StringVar(&libvirtcfg.PoolName, "pool-name", defaultPoolName, "libvirt storage pool")
	flags.StringVar(&libvirtcfg.NetworkName, "network-name", defaultNetworkName, "libvirt network pool")
	flags.StringVar(&libvirtcfg.DataDir, "data-dir", defaultDataDir, "libvirt storage dir")
	flags.BoolVar(&libvirtcfg.DisableCVM, "disable-cvm", false, "Use non-CVMs for peer pods")
	flags.StringVar(&libvirtcfg.LaunchSecurity, "launch-security", defaultLaunchSecurity, "Libvirt's LaunchSecurity element for Confidential VMs. SEV or s390-pv. If omitted, will automatically determine.")
	flags.StringVar(&libvirtcfg.Firmware, "firmware", defaultFirmware, "Path to OVMF")

}

func (*Manager) LoadEnv() {
	cloud.DefaultToEnv(&libvirtcfg.URI, "LIBVIRT_URI", defaultURI)
	cloud.DefaultToEnv(&libvirtcfg.PoolName, "LIBVIRT_POOL", defaultPoolName)
	cloud.DefaultToEnv(&libvirtcfg.NetworkName, "LIBVIRT_NET", defaultNetworkName)
	cloud.DefaultToEnv(&libvirtcfg.VolName, "LIBVIRT_VOL_NAME", defaultVolName)
	cloud.DefaultToEnv(&libvirtcfg.LaunchSecurity, "LIBVIRT_LAUNCH_SECURITY", defaultLaunchSecurity)
	cloud.DefaultToEnv(&libvirtcfg.Firmware, "LIBVIRT_FIRMWARE", defaultFirmware)
}

func (*Manager) NewProvider() (cloud.Provider, error) {
	return NewProvider(&libvirtcfg)
}
