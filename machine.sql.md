# 数据库文档

## machine_info表(机器信息)

| 字段名称  |  字段类型   |         字段含义         |
|-----------|-------------|--------------------------|
| id        | int(11)     | 自增id                   |
| name      | VARCHAR(40) | 机器的名字               |
| resume    | VARCHAR(50) | 机器的简介               |
| type      | VARCHAR(10) | 机器类型                 |
| site      | int(11)     | 机位                     |
| available | BOOLEAN     | 设备是否可用             |
| bind      | VARCHAR(40) | 绑定的标识，比如AGENT ID |
| ctime     | int         | 创建时间                 |
## order_info表(订单的信息)

| 字段名称 |  字段类型   |                       字段含义                       |
|----------|-------------|------------------------------------------------------|
| id       | int(11)     | 自增id                                               |
| mid      | int(11)     | 对应的machine id                                     |
| status   | VARCHAR(40) | 该订单是否cancel、done、going、wait，只能有一个going |
| ctime    | int         | 创建时间                                             |
## class_info表(订单的信息)

| 字段名称  |  字段类型   |     字段含义     |
|-----------|-------------|------------------|
| id        | int(11)     | 自增id           |
| mid       | int(11)     | 对应的machine id |
| name      | VARCHAR(25) | 节点的名字       |
| resume    | VARCHAR(50) | 机器的简介       |
| number    | int(11)     | 节点编号         |
| available | BOOLEAN     | 设备是否可用     |
| ctime     | int         | 创建时间         |
