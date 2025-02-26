package jieba

import (
	"fmt"
	"github.com/ssp97/ZeroBot-Plugin/pkg/jieba"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
	"time"
)

func init() {
	zero.OnRegex(`^jieba分词\s(.+?)$`).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {

			now := time.Now()

			data := ctx.State["regex_matched"].([]string)[1]
			resWords := jieba.Seg.Cut(data, true)
			ctx.SendChain(message.Text(fmt.Sprintf("%s\t精确模式：%s \n", data, strings.Join(resWords, "-"))),
				message.Text(fmt.Sprintf("\ntime:%v",time.Since(now))))
		})


	//zero.OnRegex(`^sego分词\s(.+?)$`).SetBlock(true).SecondPriority().
	//	Handle(func(ctx *zero.Ctx) {
	//		now := time.Now()
	//		if segmenter.Dictionary() == nil{
	//			ctx.SendChain(message.Text("正在加载词典，请稍等"))
	//			segmenter.LoadDictionary("data/dict/dictionary.txt")
	//		}
	//		data := ctx.State["regex_matched"].([]string)[1]
	//		segments := segmenter.Segment([]byte(data))
	//
	//		ctx.SendChain(message.Text(sego.SegmentsToString(segments, false)),
	//			message.Text(fmt.Sprintf("\ntime:%v",time.Since(now))))
	//	})
}
