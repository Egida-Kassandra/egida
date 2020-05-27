import requests

class WorkerAPIClient:

    def __init__(self, host='localhost'):
        self.host = host

    def get_lynis_score(self):
        url = 'http://'+self.host+':3000/api/lynis'
        req = requests.get(url=url)
        if req.status_code == 200:
            return req.text
        else:
            raise Exception(req.text)

    def get_info(self):
        url = 'http://'+self.host+':3000/api/info'
        req = requests.get(url=url)
        if req.status_code == 200:
            return req.text
        else:
            raise Exception(req.text)