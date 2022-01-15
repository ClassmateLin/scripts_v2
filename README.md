# scripts_v2

## 目录说明
```text
├── Dockerfile  //  docker build
├── LICENSE
├── README.md
├── com  // 通用包
│   └── file.go
├── config 
│   ├── config.go    // 根配置
│   ├── config.json  // 配置文件
│   └── jd  // 子配置项
├── contracts  // interface 定义
│   ├── jd.go
│   └── notify.go
├── database  // database 相关
│   └── database.go
├── global  // 全局变量
│   ├── const.go
│   ├── logger.go
│   └── user_agent.go
├── go.mod
├── go.sum
├── manage.go
├── models  // 模型定义
├── repositories  // 模型操作
├── scripts // 脚本存放路径
│   └── jd  // 存放某东脚本
├── shell  // 存放shell脚本
│   └── docker-entrypoint.sh
├── storage  // 存放资源，日志文件
│   └── logs  // 日志目录
├── structs // 结构体定义
│   ├── jd.go
│   └── notify.go
└── tools // 工具函数
    ├── build_scripts.go
    ├── update_config.go
    ├── update_crontab.go
    └── update_readme.go
```

## 安装

- 拉取镜像: `docker pull classmatelin/scripts_v2:latest`
- 创建容器: `docker run -itd --name scripts_v2 classmatelin/scripts_v2:latest`
- 进入容器: `docker exec -it scripts_v2 bash`

## 使用

### 配置文件

- 配置文件: `config/config.json`:
```json
{
  "logger": {
    "directory": "storage/logs", // 日志文件夹
    "filename": "logger.log" // 日志名称
  },
  "jd": {
    "cookies": [
      "pt_pin=jd1;pt_key=sssss;",  // 最简配置
      "pt_pin=jd2;ws_key=sfsfsf;",  // 使用ws_key;
      "pt_pin=jd3:ws_key=sfsfsfsf;remark=账号3;", // remark为备注账号名
      "pt_pin=jd4;pt_key=sfasfafaf;remark=账号4;push_plus=fasfasfasfafaf;", // 备注账号4, 并且单独配置push_plus通知.
      // 同时单独配置server_jd, push_plus, tg通知
      "pt_pin=j;pt_key=s;remark=g;push_plus=af;server_j=sfsaf;tg_bot_token=sfasf;tg_user_id=2;",
      // notify_type, 1: 开启单独通知并且推送到总通知, 2 开启单独通知但不推送到总通知。
      "pt_pin=j;pt_key=s;remark=g;push_plus=af;notify_type=2",
    ]
  },
  "notify": {
    "server_j": "", // server 酱通知send key
    "push_plus": "",  // push+通知 token
    "tg": { 
      "bot_token": "",  // tg机器人token
      "user_id": ""   // tg chat id.
    }
  }
}
```


### 定时任务

- 定时任务配置文件: `config/crontab.json`
- 键为脚本名称(不带后缀.go), 值为定时任务值, 设为`false`则关闭脚本定时任务。
```json
{
  "jd_check_cookies": "0 */4 * * *",  
  "jd_joy_park": false 
}
```

### 手动执行

- `go/src/scripts`目录下的*.bin结尾的文件均为可执行文件。
- 例如: `./jd_check_cookies.bin`执行检测cookies。


### 其他

<img src="https://classmatelin.top/upload/2022/01/571641820763_.pic-7b8e58ed85294d7caf606decfbd8bbce.jpg" width="400" height="400">


## 免责声明


**一切下载及使用脚本(scripts_v2)时均被视为已经仔细阅读并完全同意以下条款**：

- 脚本(scripts_v2)仅供个人学习与交流使用，严禁用于商业以及不良用途。
- 如有发现任何商业行为以及不良用途，脚本(scripts_v2)作者有权撤销使用权。
- 使用本脚本所存在的风险将完全由其本人承担，脚本(scripts_v2)作者不承担任何责任。
- 脚本(scripts_v2)注明之服务条款外，其它因不当使用本脚本而导致的任何意外、疏忽、合约毁坏、诽谤、版权或其他知识产权侵犯及其所造成的任何损失，本脚本作者概不负责，亦不承担任何法律责任。
- 对于因不可抗力或因黑客攻击、通讯线路中断等不能控制的原因造成的服务中断或其他缺陷，导致用户不能正常使用，脚本(scripts_v2)作者不承担任何责任，但将尽力减少因此给用户造成的损失或影响。
- 本声明未涉及的问题请参见国家有关法律法规，当本声明与国家有关法律法规冲突时，以国家法律法规为准。 
- 本脚本相关声明版权及其修改权、更新权和最终解释权均属脚本(scripts_v2)作者所有。
- 您必须在下载后的24小时内从计算机或手机中完全删除以上内容。
