[[中文]](readme.md) [[English]](readme_en.md)

# 使用方法

API 文档
* https://service.autoxing.com/docs/api/zh-cn/
* https://service.autoxing.com/docs/api/en-us/
* https://serviceglobal.autoxing.com/docs/api/zh-cn/
* https://serviceglobal.autoxing.com/docs/api/en-us/



## 修改配置文件

config.py

``` python

修改内容为获取到的对应值
config["APPID"]
config["APPSecret"]
config["Authorization"] 

```

``` python
config["URLPrefix"] = 
中国区用户 "https://api.autoxing.com"
海外区用户 "https://apiglobal.autoxing.com"

```

``` python
config["token"] = AxToken.py 中获取到的token

```




## 如何获取token

AxToken.py

``` python
    manager = TokenManager()
    print(manager.getToken())
```

## 如何获取机器人列表

AxRobot.py

``` python
    manager = RobotManager(config["token"])
    ok,robots = manager.getRobotList()

```

## 如何获取POI信息

AxMapInfo.py

getPoiList 可以根据需要设置参数。businessId,robotId,areaId
一般开发，推荐robotId



``` python

        manager = MapInfoManager(config["token"])
        print(manager.getPoiList('<businessId>',None,None))

        ok,pois = manager.getPoiList(None,'<robotId>',None)
        if ok:
            for item in pois:
                print(item)
        else:
            print("Get Map List Failed")

```

## 如何获取 businessId 

参考 ：
    AxBusiness.py
    AxBuilding.py


## 如何创建和执行任务

参考：
    AxTask.py


    任务的一般流程：
        1. 构造任务对象

        2. 添加任务点:
            a.构造任务点对象
            b.添加任务点动作

        3. 设置返回点：
            a.构造任务点对象
            b.添加任务点动作
        
        4. 请求接口创建任务：
            返回任务ID

        5. 请求接口执行任务：
            返回任务结果

        6. 请求接口查询任务状态
            返回任务详情（注意：非实时）

            实时状态需要通过websocket接口获取
