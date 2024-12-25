"""
    Task Management Class
"""
import json
import requests


from config import config



class Action:
    """
    Action Class
    """
    @staticmethod
    def PauseAction(duration):
        """
        Pause Action
        """
        return {
            "type": 18,
            "data": {"pauseTime": duration},
        }
    
    @staticmethod
    def PlayAudioAction(audioId):
        """
        Pause Action
        """
        return {
            "type": 5,
            "data": {
                    "mode": 1,
                    "url":"",
                    "audioId": audioId,
                    "interval": -1,
                    "num": 1,
                    "volume": 100,
                    "channel": 1,
                    "duration":-1,
                     },
        }
    

    @staticmethod
    def WaitAction(userData):
        """
        Pause Action
        """
        return {
            "type": 40,
            "data": {"userData":userData},
        }
    


class TaskPoint:
    """
    Task Point Class
    """
    def __init__(self,poi,ignoreYaw=True) -> None:

        self.pt = {}
        self.pt["areaId"] = poi["areaId"]
        self.pt["x"] = poi["coordinate"][0]
        self.pt["y"] = poi["coordinate"][1]

        if not ignoreYaw:
            self.pt["yaw"] = poi["yaw"]


        self.pt["type"] = 0
        self.pt["stopRadius"] = 1
        self.pt["ext"] = {
            "name": poi["name"],
        }
        self.pt["stepActs"]=[]

    def addStepActs(self,stepAct):
        """
        Add Step Acts
        """
        self.pt["stepActs"].append({**stepAct})

        return self




class TaskBuilder:
    """
    Task Builder Class
    """
    def __init__(self,name :str,robotId :str) -> None:

        
        self.task = {}
        self.task["name"] = name
        self.task["robotId"] = robotId
        self.task["routeMode"] = 1
        self.task["runMode"] = 1
        self.task["runNum"] = 1
        self.task["taskType"] = 4
        self.task["runType"] = 21
        self.task["sourceType"] = 6
        self.task["ignorePublicSite"] = False
        self.task["speed"] = 1.0
        self.task["taskPts"] = []


    def addTaskPt(self,tp):
        """
        Add Task Point
        """
        self.task["taskPts"].append({**tp.pt})

        return self
    
    def setBackPt(self,pt):
        """
        Add Task Point
        """
        self.task["backPt"]=pt.pt

        return self
        

    def getTask(self):
        """
        Get Task
        """
        return self.task

    



class TaskManager:
    """
    Task Management Class
    """
    def __init__(self,token) -> None:
        self.token = token


    def getTaskInfo(self,taskId):
        """
        Get Task Info
        """
    
        url = config["URLPrefix"]+f"/task/v1.1/{taskId}"

        try:
            r = requests.get(
                url, headers={"X-Token": self.token}, timeout=5)

            if r.status_code == 200:
                retData = json.loads(r.text)

                #print(retData["data"])

                if retData["status"] == 200:
                    return True,retData["data"]
                
        except requests.RequestException as e:
            print(e)


        return False,None
    


    def executeTask(self,taskId):
        """
        Execute Task
        """
    
        url = config["URLPrefix"]+f"/task/v1.1/{taskId}/execute"

        try:
            r = requests.post(
                url, headers={"X-Token": self.token}, timeout=5)

            if r.status_code == 200:
                retData = json.loads(r.text)
                if retData["status"] == 200:
                    return True
                
        except requests.RequestException as e:
            print(e)


        return False


    def newTask(self,taskData):
        """
        New Task
        """

        url = config["URLPrefix"]+"/task/v1.1"

        data = taskData

       

        try:
            r = requests.post(
                url, headers={"X-Token": self.token}, json=data, timeout=5)

            if r.status_code == 200:
                retData = json.loads(r.text)

                print(retData)

                if retData["status"] == 200:
                    return True, retData["data"]["taskId"]
                
        except requests.RequestException as e:
            print(e)


        return False, None
    

if __name__ == "__main__":

    import time

        # 任务的一般流程：
        # 1. 构造任务对象

        # 2. 添加任务点:
        #     a.构造任务点对象
        #     b.添加任务点动作

        # 3. 设置返回点：
        #     a.构造任务点对象
        #     b.添加任务点动作
        
        # 4. 请求接口创建任务：
        #     返回任务ID

        # 5. 请求接口执行任务：
        #     返回任务结果

        # 6. 请求接口查询任务状态
        #     返回任务详情（注意：非实时）

        #     实时状态需要通过websocket接口获取


    
    
    poi1 = {'areaId': '66ea87fe6cb0037e92ba0ac4', 'buildingId': '60a4c374059acc6c8bdff074', 'businessId': '66baf9be27a0744d055025be', 'coordinate': [-0.22222543918815063, 1.6403502840489637], 'floor': 16, 'floorName': '19', 'id': '676ba691635ae4debdc3bb8e', 'name': 'm1', 'oldFeatureId': '676ba691635ae4debdc3bb8e', 'properties': {}, 'type': 11, 'version': 'v23.12.14', 'yaw': 0}

    poi2 = {'areaId': '66ea87fe6cb0037e92ba0ac4', 'buildingId': '60a4c374059acc6c8bdff074', 'businessId': '66baf9be27a0744d055025be', 'coordinate': [-0.16790582975545476, 3.853874768537935], 'floor': 16, 'floorName': '19', 'id': '676ba6a1671a6202a258f696', 'name': 'm2', 'oldFeatureId': '676ba6a1671a6202a258f696', 'properties': {}, 'type': 11, 'version': 'v23.12.14', 'yaw': 0}


    task = TaskBuilder("Task1","8981307a02163yT")


    tp1 = TaskPoint(poi1)
    task.addTaskPt(tp1)


    tp2 = TaskPoint(poi2)
    tp2.addStepActs(Action.PlayAudioAction("3111002")).addStepActs(Action.PauseAction(10)).addStepActs(Action.PlayAudioAction("3111012"))
    task.addTaskPt(tp2)


    park = TaskPoint(poi1)
    park.addStepActs(Action.WaitAction({"cmd":"test"}))
    task.setBackPt(park)


    print(task.getTask())



    manager = TaskManager(config["token"])

    ok,taskID = manager.newTask(task.getTask())
    print(ok,taskID)
    if ok:
        ok = manager.executeTask(taskID)
        if ok:
            while True:
                time.sleep(1)
                ok,data = manager.getTaskInfo(taskID)
                if ok:
                    print(f'isCancel:{data["isCancel"]} isFinish:{data["isFinish"]} isExcute:{data["isExcute"]}')
                else:
                    break

            




    

    
