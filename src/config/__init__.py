def config(args):
    if args.group is not None:
        with open("/etc/ansible/hosts", "w") as f:
            f.write("[{}]".format(args.group))
            for host in args.hosts:
                f.write("{} ansible_ssh_user=root".format(host))
        with open("../common/hosts", "w") as f:
            f.write(args.group)