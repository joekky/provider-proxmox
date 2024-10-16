package config

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("proxmox_virtual_environment_vm", func(r *config.Resource) {
        r.ShortGroup = "compute"
        r.Kind = "VirtualMachine"
    })

    p.AddResourceConfigurator("proxmox_virtual_environment_container", func(r *config.Resource) {
        r.ShortGroup = "compute"
        r.Kind = "Container"
    })

    // Add more resource configurations here
}
