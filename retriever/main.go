package main

import (
	"fmt"
	"mooc_Go_study/retriever/mock"
	real2 "mooc_Go_study/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}
func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "gwf",
			"course": "golang",
		},
	)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake imooc.com"}
	inspect(&retriever)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	realretriever := r.(*real2.Retriever)
	fmt.Println(realretriever.TimeOut)

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
