package main

import "log"

// Event 事件
type Event struct {
	Info    string
	NewName string
	Ip      string
	Host    string
	Article string
	Port    int
}

type Ob interface {
	Update(e Event)
	PublishArticleNotify(e Event)
}

type Notifier interface {
	Register(ob Ob)
	Deregister(ob Ob)
	NotifyAll(event Event) bool
	PublishArticle(articleName string)
	PublishArticleNotify(event Event)
}

//FamousPerson  名人
type FamousPerson struct {
	SubPersonList []Ob
	ArticleList   []string
}

func NewObserver() FamousPerson {
	return FamousPerson{SubPersonList: make([]Ob, 0)}
}

func (this *FamousPerson) PublishArticle(articleName string) {
	this.ArticleList = append(this.ArticleList, articleName)
	this.PublishArticleNotify(Event{
		Info:    "",
		NewName: "",
		Ip:      "",
		Host:    "",
		Article: articleName,
		Port:    0,
	})
}

func (this *FamousPerson) PublishArticleNotify(event Event) {
	for i := range this.SubPersonList {
		this.SubPersonList[i].PublishArticleNotify(event)
	}
}
func (this *FamousPerson) NotifyAll(event Event) bool {
	for i := range this.SubPersonList {
		this.SubPersonList[i].Update(event)
	}
	return true
}

func (this *FamousPerson) Register(ob Ob) {
	this.SubPersonList = append(this.SubPersonList, ob)
}

func (this *FamousPerson) Deregister(ob Ob) {

}

// NormalPerson 普通人
type NormalPerson struct {
	Name string
	Ip   string
	Host string
	Port string
}

func (o *NormalPerson) Update(e Event) {
	log.Println("A.ob.Info: ", e.Info)
}

func (o *NormalPerson) PublishArticleNotify(e Event) {
	log.Println("A.PublishArticleNotify: ", e.Article)
}

// NormalPersonB 普通人B
type NormalPersonB struct {
	Name string
	Ip   string
	Host string
	Port string
}

func (o *NormalPersonB) Update(ob Event) {
	log.Println("B.ob.Info: ", ob.Info)
	o.Name = ob.NewName
}

func (o *NormalPersonB) PublishArticleNotify(e Event) {
	log.Println("B.PublishArticleNotify: ", e.Article)
}

func main() {
	observer := NewObserver()
	oA := &NormalPerson{}
	oB := &NormalPerson{
		Name: "old",
	}
	observer.Register(oA)
	observer.NotifyAll(Event{Info: "1"})
	observer.Register(oB)
	observer.NotifyAll(Event{Info: "2", NewName: "new"})
	log.Println("new.B.Name: ", oB.Name)
	observer.PublishArticle("newArticle1")
}
