# proto

## 目录结构
proto
 |- common.proto 协议公共定义
 |- retcode.proto 返回码定义
 |- xxx_server 服务协议目录
     |- xxx_service.proto 服务协议定义
     |- ......

## 协议约定
1. 后台服务接口约定
(1) 接口原型: 
    rpc XXX(XXXRequest) returns (XXXResponse)
(2) 命名约定:
    请求结构体和响应结构体分别在接口名后增加Request和Response
(3) 结构体定义:
    a. Request和Response结构体分别由header和data两部分组成
    message XXXRequest {
      // Header field
      "header-field1": xxx,
      "header-field2": xxx,
      // Data field
      XXXRequestData data = 2;
    };
    b. data字段为空时可省略该字段
2. 前端接口约定
前端使用形如以下描述的HTTP请求访问后台接口
(1) URL: /cgi-bin/version/project/server/service/fucntion, 其中
    version - 格式为vN，例如v1，表示接口版本，可用于区分接入层
    project - 项目名称，proto文件中package字段的前半部分
    server - 服务名，proto文件中package字段的后半部分
    service - 对应proto文件中的service名称
    function - 对应proto文件中service里定义的接口
(2) Content: JSON字符串，header-field对应接口中的header参数内容，data对应request中data字段
    {
        "header-field1": xxx,
        "header-field2": xxx,
        ......
        "data": {
          ......
        }
    }
(3) Response: JSON字符串，顶层json为header，内嵌data字段对应response
    {
        "retcode" : int,
        "retmsg" : string,
        "requestId": string,
        ......
        "data": {
            ......
        }
    }

