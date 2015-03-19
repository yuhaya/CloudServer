# APP与服务端接口

## 用户管理

### 用户登录

#### 请求说明

http请求方式：`post`

	http://api.domain.com/user/tickets
	
请求包结构体为:

	{"type": 0, "phone": "1821821922", "username": "zero", "password": "123456"}

#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
type | 是  | App登录发送的请求类型，0表示手机号登录，1表示用户名和密码登录
phone | 否  | 用户登录的手机号
username | 否 | 用户登录的用户名，默认为手机号码
password | 否 | 用户登录的密码

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {
			"token":"623ca3c5-254e-4dc7-a916-f436707d8de0",
			"expires":1415346248,
			"userinfo":{
				"id":"xxx",
				"username":"xxx",
				"id_card":"xxx",
				"school_name":"xxx",
				"school_id":"xx",
				"phone":"xxx",
				"gender":0/1,//0->男 1->女
				"name":"xx"//真实姓名
				}
			}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据
token | 返回令牌
expires | 令牌到期时间
userinfo | 用户相关信息


### 添加学生信息

#### 请求说明

	当前登录用户添加的学生，默认为该学生的主家长

http请求方式：`post`

	http://api.domain.com/student/create?token=TOKEN
	
请求包结构体为:

	{
		user_id:"xxxx",// 添加者（家长）的id
		name:"xxx",
		gender:0/1,
		school_id:"xxx",
		card:"xxx",
		grade:"xxx",
		class_name:"xx",
		enrol_time:"2010-09-10",
		relation:"xx"
	}

#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
user_id | 是 | 添加家长的id
name | 是  | 用户户名称
gender | 否  | 用户性别 0=》女 1=》男
card | 否 | 卡号
school_id | 是 | 学校id
grade | 否 | 年级信息
class_name | 否 | 班级信息
enrol_time | 否 | 入学时间
relation | 否 | 与添加家长的关系

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 编辑学生信息

#### 请求说明

	为该学生的主家长才可以编辑学生

http请求方式：`post`

	http://api.domain.com/student/update?token=TOKEN
	
请求包结构体为:

	{
		id:"xxxx",
		user_id:"xxx",
		name:"xxx",
		gender:0/1,
		school_id:"xx",
		card:"xxx",
		grade:"xxx",
		class_name:"xx",
		enrol_time:"2010-09-10"
	}

#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
user_id | 是 | 家长的id
id | 是 | 学生的id
name | 否  | 用户户名称
gender | 否  | 用户性别 0=》女 1=》男
card | 否 | 卡号
school_id | 否 | 学校id
grade | 否 | 年级信息
class_name | 否 | 班级信息
enrol_time | 否 | 入学时间

#### 返回结果

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 删除学生

#### 请求说明

	当前登录用户为该学生的主家长才可以删除

http请求方式：`post`

	http://api.domain.com/student/delete?token=TOKEN
	
请求包结构体为:

	{
		student_id:"xxx"
	}

#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
student_id | 是  | 学生的id

#### 返回结果

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 获取学生列表

#### 请求说明

	获得当前登录家长所有看管的学生

http请求方式： `post`
	http://api.domain.com/student/list?token=TOKEN
请求包结构体为:

	{
		user_id:"xxx"
	}	
	
	
#### 返回说明

	{
		"code":0,
		"msg":"success",
		"data":{
			{
				id:"xxxx",
				identity:0/1
				name:"xxx",
				gender:0/1,
				school_id:"xx",
				school_name:"xxx",
				card:"xxx",
				grade:"xxx",
				class_name:"xx",
				enrol_time:"2010-09-10",
				relation:"xx"
			},
			{
				id:"xxxx",
				identity:0/1
				name:"xxx",
				gender:0/1,
				school_id:"xx",
				school_name:"xxx",
				card:"xxx",
				grade:"xxx",
				class_name:"xx",
				enrol_time:"2010-09-10",
				relation:"xx"
			},
			//....
		}
	}
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据
id| 学生的id
identity | 1=》表明学生可以编辑 0=》表示不可以编辑
name   | 用户户名称
gender   | 用户性别 0=》女 1=》男
card  | 用户的卡号
school_id | 学校的id
school_name  | 学校名称
grade  | 年级信息
class_name  | 班级信息
enrol_time  | 入学时间
relation | 该家长与学生的关系


### 获取学校列表

#### 请求说明

http请求方式： `get`
	http://api.domain.com/school/list?token=TOKEN

	
#### 返回说明

	{
		"code":0,
		"msg":"success",
		"data":{
			{
				id:"xxxx",
				name:"xxx"
			},
			{
				id:"xxxx",
				name:"xxx"
			},
			//....
		}
	}
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据
id| 学校的id
name | 学校的名称


### 更新/添加 学生的相关家长信息关系

#### 请求说明

	只有该登陆家长为学生的主家长时才可以修改
	
http请求方式： `post`

	http://api.domain.com/relation/update?token=TOKEN

请求包结构体为:

	{
		student_id："xxx",
		uid:"xxx",
		relation:"xxx"
	}
	
#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
student_id | 是  | 学生的id
uid | 是 | 家长的id （不一定是该登陆家长的id，可能是其他副家长的id）
relation | 是 | 家长与学生的关系

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 删除学生与家长之间的关系

#### 请求说明

	只有该登陆家长为学生的主家长时才可以修改
	
http请求方式： `post`

	http://api.domain.com/relation/delete?token=TOKEN

请求包结构体为:

	{
		student_id："xxx",
		uid:"xxx"
	}
	
#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
student_id | 是  | 学生的id
uid | 是 | 家长的id （不一定是该登陆家长的id，可能是其他副家长的id）

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 添加用户

#### 请求说明
	
	只有登录用户为主家长才可以添加副家长

http请求方式：`post`
	
	http://api.domain.com/user/create?token=TOKEN
	
请求包结构体为:

	{
		"par_user_id": 1,    	"username": "1821821922",    	"password": "123456",    	"phone": "1821821922",    	"id_card": "421003198008050171",    	"gender": 0    	"name": "zero"
	}

#### 参数说明
	
参数 | 必须 | 说明
------------ | ------------- | ------------
par_user_id|是|主用户编号，若此编号为空则添加为主家长
username | 否 | 用户名
password | 是 | 登录密码
id_card | 否 | 用户身份证编号
gender | 否 | 0、1
name | 否 | 真实姓名

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 修改家长信息

#### 请求说明
	
	主家长可以修改自己信息以及其添加的副家长信息，或者副家长修改自己的信息

http请求方式：`post`
	
	http://api.domain.com/user/update?token=TOKEN
	
请求包结构体为:

	{
		"id": 1,    	"username": "1821821922",    	"password": "123456",    	"phone": "1821821922",    	"id_card": "421003198008050171",    	"gender": 0    	"name": "zero"
	}

#### 参数说明
	
参数 | 必须 | 说明
------------ | ------------- | ------------
id | 是 | 家长的id
username | 否 | 用户名
password | 否 | 登录密码
id_card | 否 | 用户身份证编号
gender | 否 | 0、1
name | 否 | 真实姓名

#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据


### 删除副家长

#### 请求说明
	
	删除自己添加的副家长

http请求方式：`post`
	
	http://api.domain.com/user/delete?token=TOKEN
	
请求包结构体为:

	{
		"par_user_id": "xxxxx",    	"vice_user_id": "xxxx"
	}

#### 参数说明
	
参数 | 必须 | 说明
------------ | ------------- | ------------
par_user_id | 是 | 主家长的id
vice_user_id | 是 | 副家长的id


#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 获得主家长所有添加的副家长的信息

#### 请求说明
	
	只有登录用户为主家长才可以获取

http请求方式：`get`
	
	http://api.domain.com/user/get_vice_user?token=TOKEN&id=xxxxxx

{id} 主家长
	
#### data数据返回说明
	
	{
		{
			"uid":"xxx",
			"username": "1821821922",    		"phone": "1821821922",    		"gender": 0    		"name": "zero"
		},
		{
			"uid":"xxx",
			"username": "1821821922",    		"phone": "1821821922",    		"gender": 0    		"name": "zero"
		},
		//....
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据	
uid | 家长的id
username  | 用户名
gender  | 0、1
name  | 真实姓名


## 信息管理

### 获得到校推送信息

#### 请求说明

http请求方式：`get`
	
	http://api.domain.com/message/list?token=TOKEN&uid={uid}
	
{uid} 用户的id		
	
#### data数据返回说明

	{
		{
			"id":"xxx",
			"phone":"13818151804",
			"student_id":101,
			"name":"zero",
			"gender":0,
			"school_id":"101",
			"school_name":"上海交大",
			"grade":"一年级",
			"class_name":"文科班",
			"card":"5463826199",
			"pass_time":"2014-12-11 09:35:18",
			"device_id":"gdjsjds",
			"description":"hek",
			"status":'y'
		},
		{
			"id":"xxx",
			"phone":"13818151804",
			"student_id":101,
			"name":"zero",
			"gender":0,
			"school_id":"101",
			"school_name":"上海交大",
			"grade":"一年级",
			"class_name":"文科班",
			"card":"5463826199",
			"pass_time":"2014-12-11 09:35:18",
			"device_id":"gdjsjds",
			"description":"hek",
			"status":'y'
		},
		//....
	}

参数 |  说明
------------ | ------------
id | 消息的唯一标示
phone | 手机号
student_id | 学校id
name | 姓名
gender | 性别 0-》男 1-》女
school_id | 学校id
school_name | 学校名称
grade| 年级
class_name| 班级
card| 卡号
pass_time| 时间
device_id| 设备标示
description| 设备描述
status| 设备状态 y 正常 n异常

### 通知服务端信息已接收接口 

#### 请求说明

	接受消息后，将接收消息状态改变为已接收

http请求方式：`post`
	
	http://api.domain.com/message/flag?token=TOKEN
	
请求包结构体为:

	{
		"id": "xxx"
	}

#### 参数说明
	
参数 | 必须 | 说明
------------ | ------------- | ------------
id | 是 | 消息的唯一标示


#### 返回说明

	{
		"code": 0,
		"msg": "success",
		"data": {}
	}
	
参数 |  说明
------------ | ------------- | ------------
code |  错误代码，0代代表没有错误
msg |  描述信息
data | 响应数据

### 获得到校时间点图片信息

#### 请求说明
	

http请求方式：`post`
	
	http://api.domain.com/image?token=TOKEN
	
请求包结构体为:

	{
		"school_id": "123",    	"time": "2014-12-10 12:10:30"
	}

#### data数据返回说明

	{
		"http://img.domain.com/xxxxx.png",
		"http://img.domain.com/xxxxx.png",
		"http://img.domain.com/xxxxx.png",
		//.....
	}



## 信息通知

### 传送device_token与用户信息

#### 请求说明

	用户登录成功后，会传回具体的用户与设备标示信息

http请求方式：`post`
	
	http://api.domain.com/device?token=TOKEN
	
请求包结构体为:

	{
		"device_token":"1cb6d299 cb7ab6ca 44ddcc28 36872059 d78b21ab 6cf638af 900dad2d b61fe8eb",
		"user_id":"xxx"
	}
#### 参数说明

参数 | 必须 | 说明
------------ | ------------- | ------------
device_token | 是 | 设备标示
user_id | 是 | 用户id

### 信息推送其它相关说明

* node.js 现有的实现的npm包，地址：https://github.com/logicalparadox/apnagent-playground，需要平叔将其集成到express中

* 该模块包中需要的证书文件`pfx.p12`由IOS端给出，（附件中）

* 服务端的主要工作

	1.接受客户端每次登陆后传回的设备标示，并关联到特定登录用户

	2.设备有新的信息发送到服务端时，服务端先匹配都特定用户，然后再通过用户找到关联的设备标示，将信息推送到APNS服务器
	
	3.服务端需要过一段时间从apns服务器上获取有效的设备标示（因为当部分用户安装app后再卸载，那它的设备标示便无效了。注意：一个ios客户端的设备标示并不是固定的）。然后比对自己服务器上的设备标示信息，删除无效的设备标示
	
	
关于node.js实现推送的包的使用，（该包中封装了对APNS服务端SOCKET通信的低级接口，实现更加方便进行信息推送，与服务端设备标示的获取。本地测试均正常）
