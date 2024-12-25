"""
    Business Management Class
"""
import json
import requests

from config import config


class BusinessManager:
    """
    Business Management Class
    """
    def __init__(self,token) -> None:
        self.token = token

    def getBusinessList(self):
        """
        Get Business List
        """
        url = config["URLPrefix"]+"/business/v1.1/list"


        try:
            r = requests.post(
                url, headers={"X-Token": self.token}, timeout=5)

            if r.status_code == 200:
                retData = json.loads(r.text)
 
                if retData["status"] == 200:
                    return True, retData["data"]["lists"]
                
        except requests.RequestException as e:
            print(e)


        return False, None
    

if __name__ == "__main__":

    manager = BusinessManager(config["token"])
    print(manager.getBusinessList())

