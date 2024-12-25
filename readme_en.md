[[中文]](readme.md) [[English]](readme_en.md)

# Usage Instructions

API DOC
* https://service.autoxing.com/docs/api/zh-cn/
* https://service.autoxing.com/docs/api/en-us/
* https://serviceglobal.autoxing.com/docs/api/zh-cn/
* https://serviceglobal.autoxing.com/docs/api/en-us/

## Modify the Configuration File

- [config.py](python3/config.py)

```python
Modify the following values to the corresponding obtained values:
config["APPID"]
config["APPSecret"]
config["Authorization"]
```

```python
config["URLPrefix"] = 
For users in China: "https://api.autoxing.com"
For overseas users: "https://apiglobal.autoxing.com"
```

```python
config["token"] = The token obtained from AxToken.py
```

## How to Obtain a Token

- [AxToken.py](python3/AxToken.py)

```python
    manager = TokenManager()
    print(manager.getToken())
```

## How to Get the Robot List

- [AxRobot.py](python3/AxRobot.py)

```python
    manager = RobotManager(config["token"])
    ok, robots = manager.getRobotList()
```

## How to Get POI Information


- [AxMapInfo.py](python3/AxMapInfo.py)

The `getPoiList` method allows you to set parameters such as `businessId`, `robotId`, and `areaId`. For general development, it is recommended to use `robotId`.

```python
        manager = MapInfoManager(config["token"])
        print(manager.getPoiList('<businessId>', None, None))

        ok, pois = manager.getPoiList(None, '<robotId>', None)
        if ok:
            for item in pois:
                print(item)
        else:
            print("Get Map List Failed")
```

## How to Get `businessId`

Refer to:
- [AxBusiness.py](python3/AxBusiness.py)
- [AxBuilding.py](python3/AxBuilding.py)

## How to Create and Execute Tasks

Refer to:

- [AxTask.py](python3/AxTask.py)

### General Task Workflow:
1. Construct the task object.
2. Add task points:
   - a. Construct the task point object.
   - b. Add actions to the task point.
3. Set a return point:
   - a. Construct the task point object.
   - b. Add actions to the task point.
4. Call the API to create the task:
   - Returns the task ID.
5. Call the API to execute the task:
   - Returns the task result.
6. Call the API to query the task status:
   - Returns task details (Note: not in real-time).
   - For real-time status, use the WebSocket API.
