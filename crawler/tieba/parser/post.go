package parser

// 解析会员信息

import (
	"bytes"
	"dali.cc/ccmouse/crawler/engine"
	model2 "dali.cc/ccmouse/crawler/tieba/model"
	"dali.cc/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
	"os"
	"path"
	"regexp"
	"strings"
)

var nextPageReg = regexp.MustCompile(`<a href="(.+)">下一页</a>`)
var idReg = regexp.MustCompile(`http://tieba.baidu.com/p/(.+)\?pn=(.+)`)
var id1Reg = regexp.MustCompile(`http://tieba.baidu.com/p/(.+)`)

func visitFloor(fs []model2.Floor, n *html.Node) []model2.Floor {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && strings.Index(a.Val, "l_post j_l_post l_post_bright") != -1 { //直接得内容
				floor := model2.Floor{}
				text := visitContent(nil, n)
				floor.Content = strings.Join(text, ":")
				fs = append(fs, floor)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fs = visitFloor(fs, c)
	}
	return fs
}

func visitContent(text []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "d_post_content j_d_post_content  clearfix" { //直接得内容
				for d := n.FirstChild; d != nil; d = d.NextSibling {
					if d.Data == "img" {
						// todo 下载图片
						src := getNodeVal("src", d)
						utils.DownLoadImgToDir(src, "tieba_img")
						setNodeVal("src", "tieba_img/" + path.Base(src), d)
					}
					text = append(text, d.Data)
				}
				text = append(text, n.FirstChild.Data)
				// 找到后就可退出来
				return text
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = visitContent(text, c)
	}
	return text
}

func getNodeVal(key string, n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}
func setNodeVal(key, val string, n *html.Node) bool {
	for index, a := range n.Attr {
		if a.Key == key {
			n.Attr[index].Val = val
			return true
		}
	}
	return false
}

func ParsePost(contents []byte, url, name string) engine.ParseResult {
	rs := engine.ParseResult{}
	post := model2.Post{Title: name}
	id := extractString([]byte(url), idReg)
	if (id == "") {
		id = extractString([]byte(url), id1Reg)
	}
	doc, err := html.Parse(bytes.NewReader(contents))

	if (err != nil) {
		log.Fatal().Msgf("%s 出错了", url)
	}
	post.Floors = visitFloor(post.Floors, doc)

	f, err := os.Create(id + ".html")
	defer f.Close()
	if err != nil {
		log.Printf("%s 创建失败: %s", url, err)
	}

	err = html.Render(f, doc)
	if err != nil {
		log.Printf("%s 写入失败: %s", url, err)
	}

	item := engine.Item{
		Url:     url,
		//Payload: post,
		Type:    "tieba",
		Id:      url,
	}
	rs.Items = []engine.Item{item}
	if next := extractString(contents, nextPageReg); next != "" {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   "http://tieba.baidu.com" + extractString(contents, nextPageReg),
			Parse: NewPostParser(name),
		})
	}
	// 取下一页的链接返回
	return rs
}

func extractString(c []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		return string(bytes.Join(match[1:], []byte{0x65} ))
	} else {
		return ""
	}
}

// 生成用户解析函数的函数
//func ProfileParser(name string) engine.ParserFunc   {
//	return func(body []byte, url string) engine.ParseResult {
//		return ParseProfile(body, url, name)
//	}
//}

type PpostParser struct {
	userName string
}

func (p *PpostParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParsePost(contents, url, p.userName)
}

func (p *PpostParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewPostParser(name string) *PpostParser {
	return &PpostParser{
		userName: name,
	}
}
