package im

const _API_MSG_SEND_MSG_URL = "https://api.netease.im/nimserver/msg/sendMsg.action"

type MsgSendMsgParam struct {
	// 发送者accid，用户帐号，最大32字符，
	// 必须保证一个APP内唯一
	// 是否必须: true
	From string `json:"from"`

	// 0：点对点个人消息，1：群消息（高级群），其他返回414
	// 是否必须: true
	Ope int `json:"ope"`

	// ope==0是表示accid即用户id，ope==1表示tid即群id
	// 是否必须: true
	To string `json:"to"`

	// 0 表示文本消息,
	// 1 表示图片，
	// 2 表示语音，
	// 3 表示视频，
	// 4 表示地理位置信息，
	// 6 表示文件，
	// 10 表示提示消息，
	// 100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	// 是否必须: true
	Type int `json:"type"`

	// 最大长度5000字符，JSON格式。
	// 具体请参考： 消息格式示例
	// 是否必须: true
	Body string `json:"body"`

	// 对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
	// true或false, 默认false。
	// 只对消息类型为：100 自定义消息类型 的消息生效。
	// 是否必须: false
	Antispam string `json:"antispam,omitempty"`

	// 在antispam参数为true时生效。
	// 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
	// {"type":1,"data":"custom content"}
	// 字段说明：
	// 1. type: 1：文本，2：图片。
	// 2. data: 文本内容or图片地址。
	// 是否必须: false
	AntispamCustom string `json:"antispamCustom,omitempty"`

	// 发消息时特殊指定的行为选项,JSON格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值 ，option示例:
	// {"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}
	// 字段说明：
	// 1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
	// 2. history: 该消息是否存云端历史，默认true；
	// 3. sendersync: 该消息是否需要发送方多端同步，默认true；
	// 4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
	// 5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
	// 6. badge:该消息是否需要计入到未读计数中，默认true;
	// 7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
	// 8. persistent: 是否需要存离线消息，不设置该参数时默认true;
	// 9. sessionUpdate: 是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
	// 是否必须: false
	Option string `json:"option,omitempty"`

	// 推送文案,最长500个字符。具体参见 推送配置参数详解。
	// 是否必须: false
	Pushcontent string `json:"pushcontent,omitempty"`

	// 必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
	// 是否必须: false
	Payload string `json:"payload,omitempty"`

	// 开发者扩展字段，长度限制1024字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// 发送群消息时的强推用户列表（云信demo中用于承载被@的成员），格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
	// 是否必须: false
	Forcepushlist string `json:"forcepushlist,omitempty"`

	// 发送群消息时，针对强推列表forcepushlist中的用户，强制推送的内容
	// 是否必须: false
	Forcepushcontent string `json:"forcepushcontent,omitempty"`

	// 发送群消息时，强推列表是否为群里除发送者外的所有有效成员，true或false，默认为false
	// 是否必须: false
	Forcepushall string `json:"forcepushall,omitempty"`

	// 可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	// 是否必须: false
	Bid string `json:"bid,omitempty"`

	// 可选，单条消息是否使用易盾反垃圾，可选值为0。
	// 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	// 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	// 是否必须: false
	UseYidun int `json:"useYidun,omitempty"`

	// 可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	// 是否必须: false
	YidunAntiCheating string `json:"yidunAntiCheating,omitempty"`

	// 可选，群消息是否需要已读业务（仅对群消息有效），0:不需要，1:需要
	// 是否必须: false
	MarkRead int `json:"markRead,omitempty"`

	// 是否为好友关系才发送消息，默认否
	// 注：使用该参数需要先开通功能服务
	// 是否必须: false
	CheckFriend bool `json:"checkFriend,omitempty"`

	// 自定义消息子类型，大于0
	// 是否必须: false
	SubType int `json:"subType,omitempty"`

	// 发送方是否无感知。0-有感知，1-无感知。若无感知，则消息发送者无该消息的多端、漫游、历史记录等。
	// 是否必须: false
	MsgSenderNoSense int `json:"msgSenderNoSense,omitempty"`

	// 接受方是否无感知。0-有感知，1-无感知。若无感知，则消息接收者者无该消息的多端、漫游、历史记录等
	// 是否必须: false
	MsgReceiverNoSense int `json:"msgReceiverNoSense,omitempty"`

	// 所属环境，根据env可以配置不同的抄送地址
	// 是否必须: false
	Env string `json:"env,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E5%8F%91%E9%80%81%E6%99%AE%E9%80%9A%E6%B6%88%E6%81%AF
// 给用户或者高级群发送普通消息，包括文本，图片，语音，视频和地理位置
func (y *YunxinIM) ApiMsgSendMsg(param *MsgSendMsgParam) *ImResp {
	return y.PostFrom(_API_MSG_SEND_MSG_URL, param)
}
