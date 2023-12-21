package servers

import "github.com/gofiber/fiber/v2"

type Fiber struct {
}

func NewFiber() *Fiber {
	return &Fiber{}
}

func (f *Fiber) Init() {
	fib := fiber.New()

	fib.Use(func(c *fiber.Ctx) error {
		//do whatever you need here
		return c.Next()
	})

	fib.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("{\"status\": \"OK\"}")
	})

	testGroup := fib.Group("/test")
	testGroup.Get("/ok", func(ctx *fiber.Ctx) error {
		return ctx.SendString("{\"test\": \"OK\"}")
	})
	testGroup.Get("/ko", func(ctx *fiber.Ctx) error {
		return ctx.SendString("{\"test\": \"KO\"}")
	})

	fib.Listen(":8086")
}
