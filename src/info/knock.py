import socket
import time


def knock(host, ports, delay=200, protocol="tcp"):
    """Knock host and port using tcp connection"""

    for port in ports:
        if protocol == 'tcp':
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        else:
            sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

        sock.setblocking(False)

        try:
            sock.connect((host, port))
        except socket.error:
            # Nothing really there...
            print("Error")
            pass

        time.sleep(delay)
