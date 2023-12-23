package servers

import (
	"github.com/go-martini/martini"
)

type Martini struct {
}

func NewMartini() *Martini {
	return &Martini{}
}

func (ma *Martini) Init() {
	server := martini.Classic()

	server.Use(func() {
		//do whatever you need here
	})

	server.Get("/health", func() string {
		return "{\"status\": \"OK\"}"
	})

	server.Group("/test", func(router martini.Router) {
		router.Get("/ok", func() string {
			return "{\"test\": \"OK\"}"
		})
		router.Get("/ko", func() string {
			return "{\"test\": \"KO\"}"
		})
	})

	server.RunOnAddr(":8083")
}
