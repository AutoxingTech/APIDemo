"""
Token manager, responsible for obtaining tokens, caching tokens, and avoiding duplicate retrievals
"""
import hashlib
import time
import json
import requests

from config import config


class TokenManager:
    """
    Token manager, responsible for obtaining tokens, caching tokens, and avoiding duplicate retrievals
    """
    def __init__(self) -> None:
        self.token = None
        self.expireTime = None
        self.key = None
        self.time = 0
        self.ok = False

    def getToken(self):

        if self.ok:
            current_time = time.time()
            if current_time  < self.time/1000.0 + self.expireTime:
                return True, self.token

        return self.getTokenFromServer()

    def getTokenFromServer(self):
        url = config["URLPrefix"]+"/auth/v1.1/token"

        data = {
            "appId": config["APPID"],
            "timestamp": int(time.time()*1000),
            "sign": ""  # 签名，规则为Md5(appId+timestamp+appSecret)
        }

        data["sign"] = hashlib.md5((config["APPID"] + str(data["timestamp"]) + config["APPSecret"]).encode()).hexdigest()

        try:
            r = requests.post(
                url, headers={"Authorization": config["Authorization"]}, json=data, timeout=5)

            if r.status_code == 200:

                retData = json.loads(r.text)

                if retData["status"] == 200:
                    self.key = retData["data"]["key"]
                    self.token = retData["data"]["token"]
                    self.expireTime = retData["data"]["expireTime"]
                    self.time = data["timestamp"]
                    self.ok = True

                    return True, self.token
        except requests.RequestException as e:
            print(e)

        self.ok = False
        return False, None
    
if __name__ == "__main__":

    manager = TokenManager()
    print(manager.getToken())
