package ansible

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/antonioalfa22/go-utils/collections"
	"github.com/antonioalfa22/go-utils/command"
	"github.com/antonioalfa22/go-utils/io"
)

func CreatePlaybook(tags []string, connection string) {
	createFile(connection)
	err := command.RunCommandPrintOutput("ansible-playbook", "/etc/egida/generated.yml",
		"--tags="+getTags(tags))
	if err != nil {
		fmt.Println("Error on running playbook, Do you have Ansible installed?")
	}
}

func getTags(tags []string) string {
	return strings.Join(tags[:], ",")
}

func createFile(connection string) {
	varsfiles, hostlist := getVarsAndHosts()
	qs := crearMenu(varsfiles, hostlist)
	respuestas := struct {
		VarsFile   string
		HostsGroup string
	}{}
	// RESPUESTAS
	_ = survey.Ask(qs, &respuestas)
	vars := collections.Map(io.ReadFile(respuestas.VarsFile), func(x string) string {
		return "    " + x + "\n"
	}).([]string)
	renderFile(vars, respuestas.HostsGroup, connection)
}

func renderFile(vars []string, hosts string, connection string) {
	type Options struct {
		Hosts      string
		Connection string
		Vars       []string
	}
	options := Options{Hosts: hosts, Vars: vars, Connection: connection}
	tmpl := template.New("PlaybookTemplate")
	tmpl, _ = tmpl.Parse(
		"---\n" +
			"\n" +
			"- name: Harden Server\n" +
			"  hosts:\n" +
			"    {{ .Hosts }}\n" +
			"  connection: {{ .Connection }}\n" +
			"  become: yes\n" +
			"  vars:\n" +
			"{{ range .Vars }}{{ . }}{{ end }}\n" +
			"\n" +
			"  roles:\n" +
			"    - egida-role-cis\n")
	f, err := os.Create("/etc/egida/generated.yml")
	if err != nil {
		fmt.Println("Error creating playbook: " + err.Error())
	}
	_ = tmpl.Execute(f, options)
}

func getVarsAndHosts() ([]string, []string) {
	varsroot := "/etc/egida/vars"
	var varsfiles []string
	err := filepath.Walk(varsroot, func(path string, info os.FileInfo, err error) error {
		if path != "/etc/egida/vars" {
			varsfiles = append(varsfiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Cant find /etc/egida/vars directory")
	}
	hostlist := io.ReadFile("/etc/egida/hostsgroups")
	return varsfiles, hostlist
}

func crearMenu(varsfiles []string, hostslist []string) []*survey.Question {
	var qs = []*survey.Question{
		{
			Name: "VarsFile",
			Prompt: &survey.Select{
				Message: "Which vars file do you want to use?",
				Options: varsfiles,
				Default: "/etc/egida/vars/vars_template.yml",
			},
		},
		{
			Name: "HostsGroup",
			Prompt: &survey.Select{
				Message: "Which hosts group do you want to use?",
				Options: hostslist,
				Default: "local",
			},
		},
	}
	return qs
}
