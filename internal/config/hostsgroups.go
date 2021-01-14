package config

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/collections"
	"github.com/antonioalfa22/go-utils/io"
)

func AddHostGroup(group string, hostslist []string) {
	groups := io.ReadFile("/etc/egida/hostsgroups")
	if collections.Find(hostslist, func(s string) bool { return s == group }) == nil {
		groups = append(groups, group)
		io.WriteFile(groups, "/etc/egida/hostsgroups")
		lines := io.ReadFile("/etc/ansible/hosts")
		g := "["+group+"]"
		lines = append(lines, g)
		for _, host := range hostslist {
			lines = append(lines, host+ " ansible_ssh_user=root")
		}
		io.WriteFile(lines, "/etc/ansible/hosts")
	} else {
		fmt.Println("Group " + group + " already exists")
	}
}
