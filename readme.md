#数据库增删改查一般套路
#一、安装初始化
xorm.NewSession(driverName,dataSourceName)
#二、定义实体
模型层model或者实体层entity
##1、定义与表结构对应对象User
```go
type User struct {
    Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
    Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
    Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 什么角色
    Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
    Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`   // 什么角色
    Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 什么角色
    Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 什么角色
    Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
    Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   // 什么角色
    Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   // 什么角色
    Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
}
```
#三、定义和业务相关的服务
服务层service,专门用来存放数据库业务服务的,如
注册、登录
##2、查询单个用户Find,参数userId
       DbEngin.ID(userId).Get(&User)
##3、查询满足某一类条件的Search
       //
       result :=make([]User,0)
       DbEngin.where("mobile=? ",moile).Find(&result)
       DbEngin.where("mobile=? ",moile).Get(&User)
##4、创建一条记录Create
       DBengin.InsertOne(&User)
##5、修改某条记录Update
     DBengin.ID(userId).Update(&User)
     // update ... where id = xx
     DBengin.Where("a=? and b=?",a,b).Update(&User)
     DBengin.Where("a=? and b=?",a,b).Cols("nick_name").Update(&User)
##6、删除某条记录Delete
     DBengin.ID(userId).Delete(&User)
##7、MD5加密函数
```cgo
import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string{
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	return  hex.EncodeToString(cipherStr)

}
func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePasswd(plainpwd,salt,passwd string) bool{
	return Md5Encode(plainpwd+salt)==passwd
}
func MakePasswd(plainpwd,salt string) string{
	return Md5Encode(plainpwd+salt)
}
```     
#四、控制器层调用
```go
var userServer server.UserServer
type UserCtrl struct{}

func (ctrl *UserCtrl)Register(w){
    user = userServer.Register(mobile,passwd)
}

```

#设计可以无限扩张业务场景的消息通讯结构
```cgo
func recvproc(node *Node) {
	for{
		_,data,err := node.Conn.ReadMessage()
		if err!=nil{
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		//dispatch(data)
		fmt.Printf("recv<=%s",data)
	}
}
```
##原理
前端通过websocket发送`json格式的字符串`
用户2向用户3发送文字消息hello
```json5
{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
```
里面携带
谁发的-userid
要发给谁-dstid
这个消息有什么用-cmd
消息怎么展示-media
消息内容是什么-(url,amout,pic,content等)
##核心数据结构
```cgo
type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"` //消息ID
	//谁发的
	Userid  int64  `json:"userid,omitempty" form:"userid"` //谁发的
	//什么业务
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` //群聊还是私聊
	//发给谁
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`//对端用户ID/群ID
	//怎么展示
	Media   int    `json:"media,omitempty" form:"media"` //消息按照什么样式展示
	//内容是什么
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	//图片是什么
	Pic     string `json:"pic,omitempty" form:"pic"` //预览图片
	//连接是什么
	Url     string `json:"url,omitempty" form:"url"` //服务的URL
	//简单描述
	Memo    string `json:"memo,omitempty" form:"memo"` //简单描述
	//其他的附加数据，语音长度/红包金额
	Amount  int    `json:"amount,omitempty" form:"amount"` //其他和数字相关的
}
const (
    //点对点单聊,dstid是用户ID
	CMD_SINGLE_MSG = 10
	//群聊消息,dstid是群id
	CMD_ROOM_MSG   = 11
	//心跳消息,不处理
	CMD_HEART      = 0
	
)
const (
    //文本样式
	MEDIA_TYPE_TEXT=1
	//新闻样式,类比图文消息
	MEDIA_TYPE_News=2
	//语音样式
	MEDIA_TYPE_VOICE=3
	//图片样式
	MEDIA_TYPE_IMG=4
	
	//红包样式
	MEDIA_TYPE_REDPACKAGR=5
	//emoj表情样式
	MEDIA_TYPE_EMOJ=6
	//超链接样式
	MEDIA_TYPE_LINK=7
	//视频样式
	MEDIA_TYPE_VIDEO=8
	//名片样式
	MEDIA_TYPE_CONCAT=9
	//其他自己定义,前端做相应解析即可
	MEDIA_TYPE_UDEF=100
)
/**
消息发送结构体,点对点单聊为例
1、MEDIA_TYPE_TEXT
{id:1,userid:2,dstid:3,cmd:10,media:1,
content:"hello"}

3、MEDIA_TYPE_VOICE,amount单位秒
{id:1,userid:2,dstid:3,cmd:10,media:3,
url:"http://www.a,com/dsturl.mp3",
amount:40}

4、MEDIA_TYPE_IMG
{id:1,userid:2,dstid:3,cmd:10,media:4,
url:"http://www.baidu.com/a/log.jpg"}


2、MEDIA_TYPE_News
{id:1,userid:2,dstid:3,cmd:10,media:2,
content:"标题",
pic:"http://www.baidu.com/a/log,jpg",
url:"http://www.a,com/dsturl",
"memo":"这是描述"}


5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分
{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
6、MEDIA_TYPE_EMOJ 6
{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}

7、MEDIA_TYPE_Link 7
{id:1,userid:2,dstid:3,cmd:10,media:7,
"url":"http://www.a.com/dsturl.html"
}

8、MEDIA_TYPE_VIDEO 8
{id:1,userid:2,dstid:3,cmd:10,media:8,
pic:"http://www.baidu.com/a/log,jpg",
url:"http://www.a,com/a.mp4"
}

9、MEDIA_TYPE_CONTACT 9
{id:1,userid:2,dstid:3,cmd:10,media:9,
"content":"10086",
"pic":"http://www.baidu.com/a/avatar,jpg",
"memo":"胡大力"}

*/
```
##从哪里接收数据?怎么处理这些数据呢?
```cgo
func recvproc(node *Node) {
	for{
		_,data,err := node.Conn.ReadMessage()
		if err!=nil{
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		fmt.Printf("recv<=%s",data)
		dispatch(data)
	}
}
func dispatch(data []byte){
    //todo 转成message对象
    
    //todo 根据cmd参数处理逻辑
    
    
    
    
    
    
    msg :=Message{}
    err := json.UnMarshal(data,&msg)
    if err!=nil{
        log.Printf(err.Error())
        return ;
    }
    switch msg.Cmd {
    	case CMD_SINGLE_MSG: //如果是单对单消息,直接将消息转发出去
    		//向某个用户发回去
    		fmt.Printf("c2cmsg %d=>%d\n%s\n",msg.Userid,msg.Dstid,string(tmp))
    		SendMsgToUser(msg.Userid, msg.Dstid, tmp)
    		//fmt.Println(msg)
    	case CMD_ROOM_MSG: //群聊消息,需要知道
    		fmt.Printf("c2gmsg %d=>%d\n%s\n",msg.Userid,msg.Dstid,string(tmp))
    		SendMsgToRoom(msg.Userid, msg.Dstid, tmp)
    	case CMD_HEART:
    	default:
    	    //啥也别做
    	    
    	}
    		
}
```

