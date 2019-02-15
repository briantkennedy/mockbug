package example

type ThingApplyArg struct {
	SomeList []string
}

type Thing interface {
	Apply(a *ThingApplyArg)
}

