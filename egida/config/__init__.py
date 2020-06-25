def config(args):
    if args.group is not None:
        with open("/etc/ansible/hosts", "w") as f:
            f.write("[{}]\n".format(args.group))
            for host in args.hosts:
                f.write("{} ansible_ssh_user=root\n".format(host))
        with open("egida/common/hosts", "w") as f:
            f.write(args.group)