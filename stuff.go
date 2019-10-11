package counterfeiterrepro

//go:generate go run -mod=vendor github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . ThingDoer

type ThingDoer interface {
	DoThing() error
}
