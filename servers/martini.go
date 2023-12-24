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
	server := lightMartini()

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

func lightMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}
