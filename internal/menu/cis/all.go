package cis

import "github.com/antonioalfa22/egida/internal/ansible"

func ShowAllMenu(connection string) {
	ansible.CreatePlaybook([]string{}, connection)
}
