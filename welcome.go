package welcome

import (
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &welcome{}
var logger = utils.GetModuleLogger("logiase.autoreply")

type welcome struct {
}

func (self *welcome) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "otakuluo.welcome",
		Instance: instance,
	}
}

func (self *welcome) Init() {
}

func (self *welcome) PostInit() {
}

func (self *welcome) Serve(b *bot.Bot) {
	b.GroupMemberJoinEvent.Subscribe(self.ProcessNewMemeberJoint)
}

func (self *welcome) Start(bot *bot.Bot) {
}

func (self *welcome) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func (self *welcome) ProcessNewMemeberJoint(client *client.QQClient, event *client.MemberJoinGroupEvent) {
	responseMsg := message.NewSendingMessage().Append(message.NewAt(event.Member.Uin)).Append(message.NewText("欢迎来到部队群！进群请优先改群昵称先混眼熟哦。\r\n\r\n新人推荐优先了解以下内容：\r\n1、招待码↓（师徒结对，详询群主）\r\nhttps://actff1.web.sdo.com/20190315Zhaodai/index.html#/index\r\n2、新手入门攻略站↓（自强芽推荐）\r\nhttps://ff14.org/?utm_source=ffcafe&utm_medium=website&utm_campaign=navbar\r\n3、游戏中文维基：↓\r\nhttps://ff14.huijiwiki.com/wiki/%E9%A6%96%E9%A1%B5\r\n4、禁止【买金】【代练】【代打】，FF14游戏环境并不包容此类行为。\r\n5、下本请做好职业本职工作，及时更新装备！打本中遇到什么问题，推荐出本后群内求助，请不要在副本争吵浪费时间。\r\n6、支持讨论辩论，但请勿攻击他人，互相尊重。禁止键政内容。\r\n\r\n希望小伙伴顺利度过游戏前期，游戏愉快哦~\r\n"))
	client.SendGroupMessage(event.Group.Code, responseMsg)
}
