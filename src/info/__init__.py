from src.info.worker_api_client import WorkerAPIClient
from src.common.command import run_command
from multiprocessing import Pool


def info(args):
    hosts = args.hosts
    audit = args.audit
    for host in hosts:
        knock(host=host, ports=[7001, 8002, 9003])
    if audit:
        pool = Pool(4)
        pool.map(audit_worker, hosts)
    for x in get_info(hosts):
        print(x)
    for host in hosts:
        knock(host=host, ports=[9003, 8002, 7001])


def audit_worker(host):
    api = WorkerAPIClient(host=host)
    try:
        print("Host: ", host, ", Score: ", api.get_lynis_score())
    except:
        print("Error on connect to host: ", host)


def get_info(hosts):
    pool = Pool(4)
    result = pool.map(info_worker, hosts)
    worker_info = [x for x in result if result is not None]
    return worker_info


def info_worker(host):
    api = WorkerAPIClient(host=host)
    try:
        return api.get_info()
    except:
        print("Error on connect to host: ", host)
        return None


def knock(host, ports):
    run_command('knock ' + str(host) + ' '.join(str(port) for port in ports))
