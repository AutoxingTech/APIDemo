"""
    Map Info  Management Class
"""
import json
import requests

from config import config

class MapInfoManager:
    """
    Map Info Management Class
    """
    def __init__(self,token) -> None:
        self.token = token

    def getPoiList(self,businessId,robotId,areaId):
        """
        Get Poi List
        """

        url = config["URLPrefix"]+"/map/v1.1/poi/list"

        data = {}

        if businessId:
            data["businessId"] = businessId
        if robotId:
            data["robotId"] = robotId
        if areaId:
            data["areaId"] = areaId


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
    

if __name__ == "__main__":
    manager = MapInfoManager(config["token"])
    print(manager.getPoiList('<businessId>',None,None))
    ok,pois = manager.getPoiList(None,'<robotId>',None)
    if ok:
        for item in pois:
            print(item)
    else:
        print("Get Poi List Failed")

