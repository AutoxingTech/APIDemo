"""
    Building Management Class
"""
import json
import requests

from config import config


class BuildingManager:
    """
    Building Management Class
    """
    def __init__(self,token) -> None:
        self.token = token

    def getBuildingList(self):
        """
        Get Building List
        """
        url = config["URLPrefix"]+"/building/v1.1/list"


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

    manager = BuildingManager(config["token"])
    print(manager.getBuildingList())

