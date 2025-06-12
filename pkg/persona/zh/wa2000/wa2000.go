// Persona for WA2000 character of Girl Frontline 2.
package wa2000

import (
	"math/rand/v2"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dark-person/gf2-dc-bot/pkg/api"
	"github.com/dark-person/gf2-dc-bot/pkg/discordutils"
)

const initDefenseVal = 2

const BOT_NAME = "wa2000"

// Persona for WA2000 character.
type personaWA2000 struct {
	dialogList []string
	defense    int
}

// Interface check.
var _ api.BotPersona = (*personaWA2000)(nil)

// Create a new Persona as WA2000 character.
func New() *personaWA2000 {
	l := make([]string, 0)
	l = append(l,
		// Girl Frontline 2 Exile
		"哈……哈……總算……找到你了，快讓我上車啦！",
		"突、突然湊過來幹什麼，快走開啦……",
		"最近的財務情況不太妙？多做幾個訂單不就行了！什、沒有！我才沒關心你這邊的情況呢！",
		"不好意思和大家接觸？怎、怎麼可能！我只是單純覺得處理人際關係很麻煩！",
		"你、你想看我的槍？也不是不可以，如果是你的話……喂，一定要輕拿輕放！",
		"夜裡黃區降溫也太厲害了，看你這兩天病怏怏的，不會是凍出病來了吧？要注意保暖啊笨蛋！",
		"那個，如果你以後在黃區混不下去了，可以來綠區找我。我有個房子，面積不大，但兩個人住也夠了，我做委託賺錢——我是說如果，如果！懂不懂啊！",
		"我能不能坦率一點？我、我哪裡不坦率啦！尤其是面對你的時候！再要求更多的話……我……我做不到的……",
		"喂，喂！叫你好幾聲了都沒反應！嘶，你這黑眼圈是怎麼回事？不舒服就趕緊回去休息好吧！！",
		"這是用桑朵萊希給我的食譜做出來的甜點，你要不要嘗——才、才不是特意做給你吃的！",
		"哇，這毛絨絨的感覺……喵喵的聲音也……你、你看什麼？！我才沒有喜歡貓，是這個視頻碰巧被我看到了！",
		"對了，你、你當初為什麼不說一聲就——哼，不願意回答就算了！",
		"你這兩天怎麼跟其他人走這麼近啊？沒別的意思，就——我們不是朋友嗎，多說話才能增進感情！春、春田也是這麼說的！",
		"呆著幹什麼，還不快過來！",
		"像以前一樣，好好感謝我吧。",

		// Girl Frontline
		"你去哪了！就這麼不想看到我嗎？",
		"真是笨呢，要是真的討厭你，就不會和你說話了。",
		"聽好了！別隨便叫我的名字，還有，別隨便碰我的槍！",
		"誒？臉紅？還不是今天有點熱！別會錯意了啊！",
		"還有哪裡需要做的？讓擁有著昂貴身價的我忙了一整天，下班後該好好報答我一番吧？",
		"噫！才沒有害怕呢！哇！別拿著南瓜頭過來！",
		"反正你現在閒著吧，一個人看著太可憐了，稍微陪陪你吧……",
		"就算你不說，我也會來的，還不快感謝我。",
		"笨蛋，才不想被你擔心呢。",
		"看吧，記得感謝我哦。",
		"哼，好吧，我幹給你看。",
		"我回來了，這不是理所當然的麼？",
		"我本人都親自過來給你道謝，你就心懷感激的收下吧。",
		"我才不是在等你呢！",
		"我才沒那個意思！",
		"吶！反正你現在閒著吧……七夕遊園會可以陪你一起去哦！",
		"新年還是會好好打招呼的！新、新年快樂！怎麼了有意見嗎！",
		"湊巧多出來的巧克力啦！要是你敢說不要可不會饒了你的！",
	)

	return &personaWA2000{dialogList: l, defense: initDefenseVal}
}

func (p *personaWA2000) Help() string {
	return "" +
		"哼，記好這些指令了，沒有空格的:\n" +
		"- `?help`: 幫助訊息\n" +
		"- `!remind` : 提醒今天有什麼要做\n" +
		"- `@我` : 哼\n"
}

func (p *personaWA2000) GetReminderCustomizedStr(t time.Time) (prefix string, suffix string) {
	return "哼，是時候確認一下自己的:\n", ""
}

func (p *personaWA2000) RandomDialog() string {
	r := rand.IntN(len(p.dialogList))
	return p.dialogList[r]
}

func (p *personaWA2000) ReplyIfMentionBot(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error) {
	if strings.HasPrefix(incoming, s.State.User.Mention()) {
		err = discordutils.SendMsgToChannel(s, channelID, p.RandomDialog())
		return err != nil, err
	}

	return false, nil
}

func (p *personaWA2000) ReplyIfMentionBotName(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error) {
	lower := strings.ToLower(incoming)

	if strings.Contains(lower, BOT_NAME) {
		err = discordutils.SendMsgToChannel(s, channelID, "哼")
		return err != nil, err
	}

	return false, nil
}

func onSpecialMention() string {
	r := rand.IntN(1)
	switch r {
	case 0:
		return "(臉紅"

	case 1:
		return "(當機"

	default:
		return "(WA醬當機中)"
	}
}

func (p *personaWA2000) ReplyIfHasKeyword(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error) {
	// Special mention bot alias name
	if strings.Contains(incoming, "哇醬") || strings.Contains(incoming, "WA醬") {
		err = discordutils.SendMsgToChannel(s, channelID, onSpecialMention())
		return err != nil, err
	}

	// Burn someone
	if strings.Contains(incoming, "歐洲人都該燒") {
		err = discordutils.SendMsgToChannel(s, channelID, "(舉槍瞄準歐洲人")
		return err != nil, err
	}

	return false, nil
}
