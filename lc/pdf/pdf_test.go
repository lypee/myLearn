package pdf

import (
	"testing"
)

func TestRequestPdf_GeneratePDF(t *testing.T) {
	htmlBody := "<p>虽然我是一个单个的视频</p><p>但是我</p><p>也</p><p>是</p><p>有</p><p>笔记<span>、</span></p><p><span>表情会乱码呀～不支持呀😇<span>😇</span></span></p><table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"><tbody><tr><th>1<br/></th><th>2<br/></th><th>3<br/></th><th>4<br/></th><th>5<br/></th></tr><tr><td>3<br/></td><td><br/></td><td><br/></td><td><br/></td><td><br/></td></tr><tr><td>3<br/></td><td><br/></td><td><br/></td><td><br/></td><td><br/></td></tr><tr><td>3<br/></td><td><br/></td><td><br/></td><td><br/></td><td><br/></td></tr><tr><td>3<br/></td><td><br/></td><td><br/></td><td><br/></td><td><br/></td></tr></tbody></table><p><br/><br/></p><p><span><span></span></span></p><p>的</p><p><img src=\"https://t7.baidu.com/it/u=2604797219,1573897854&amp;fm=193&amp;f=GIF\" style=\"max-width:100%;\" contenteditable=\"false\"/><br/></p><p>哈哈哈～</p><p><br/><img src=\"https://t7.baidu.com/it/u=2604797219,1573897854&amp;fm=193&amp;f=GIF\" style=\"max-width:100%;\" contenteditable=\"false\"/></p><p><img src=\"https://t7.baidu.com/it/u=2604797219,1573897854&amp;fm=193&amp;f=GIF\" style=\"max-width:100%;\" contenteditable=\"false\"/><br/></p>"

	r := NewRequestPdf(htmlBody)

	r.GeneratePDF(htmlBody)
}
