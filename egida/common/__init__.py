# ========  Vars  ===================
connection_mode = "local"

def set_connection_mode(connection):
    global connection_mode
    connection_mode = connection

def get_connection_mode():
    global connection_mode
    return connection_mode