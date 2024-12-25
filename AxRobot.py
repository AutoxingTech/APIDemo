"""
    Robot Management Class
"""
import json
import requests

from config import config


class RobotManager:
    """
    Robot Management Class
    """
    def __init__(self,token) -> None:
        self.token = token

    def getRobotList(self):
        """
        Get Robot List
        """

        url = config["URLPrefix"]+"/robot/v1.1/list"
        data = {
            #"keyWorld": "",
            "pageSize": 10,
            "pageNum": 1
        }

        try:
            r = requests.post(
                url, headers={"X-Token": self.token}, json=data, timeout=5)

            if r.status_code == 200:
                retData = json.loads(r.text)
                if retData["status"] == 200:
                    return True, retData["data"]["list"]
                
        except requests.RequestException as e:
            print(e)


        return False, None
    

    def getRobotState(self,robotId): 
        """
        Get Robot State
        """

        url = config["URLPrefix"]+f"/robot/v1.1/{robotId}/state"
        try:
            r = requests.get(
                url, headers={"X-Token": self.token}, timeout=5)
            
            print(r.text)

            if r.status_code == 200:
                retData = json.loads(r.text)
                if retData["status"] == 200:
                    return True, retData["data"]


        except requests.RequestException as e:
            print(e)

        return False, None

if __name__ == "__main__":

    manager = RobotManager(config["token"])
    print(manager.getRobotList())
    print(manager.getRobotState("xxxxxxxxxxxx"))

    ok,robots = manager.getRobotList()
    if ok:
        for item in robots:
            print(item["robotId"],item["isOnLine"])
    else:
        print("Get Robot List Failed")
