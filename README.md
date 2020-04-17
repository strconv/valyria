# valyria
> 瓦雷利亚钢（Valyrian steel）是美国作家乔治·R·R·马丁著名的史诗奇幻小说系列《冰与火之歌》中由瓦雷利亚人发明的一种玄秘合金，能被用来制作无与伦比性能的兵器。传说在制造过程中混入了咒语和魔法，用龙焰协助锻造，使其轻便、坚韧并且从不生锈、卷刃或折断。

毕设系统的基础框架

基于 go-micro，整合 consul、go-gin、gorm、redigo等组件

封装好常用的 路由、日志、Client、dao 等模块，让开发更专注于业务逻辑

在业务开发的过程中遇到的问题，会逐渐下沉基础框架中

# Features


# 排期
- 2020.4.14
    - log、conf 100%
- 2020.4.15
    - server
        - 基础模块     100%
        - 整合 consul 100%
        - 整合 gin    100%
- 2020.4.16
    - server
        - 接入 trace 全链路跟踪 100%
        - 接入 swagger        0%
- todo
    client、gorm、redigo、jwt、自动部署