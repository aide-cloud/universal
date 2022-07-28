# universal
Basic, universal, everyone's package

## 包功能说明
- dao: 数据访问层
  - NewMysqlDao: 获取mysql连接，每次获取的连接为不同对象
  - NewMysqlDaoSingleton: 获取mysql连接，每次获取的连接为同一对象
  - NewRedisDao: 获取redis连接，每次获取的连接为不同对象
  - NewRedisDaoSingleton: 获取redis连接，每次获取的连接为同一对象

- load: 文件加载
  - Unmarshal: 获取配置文件反序列化

- executor: 服务启动模块
  - ExecMulSerProgram: 启动多个服务
  - cmd: 服务入口模块
    - New: 创建服务入口对象
  - web: web服务模块
    - New: 创建gin web服务对象
    - middleware: 中间件
      - Cross：跨域处理中间件
