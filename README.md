Ideally I should be able to:

```
go generate ./...
ginkgo
```

but counterfeiter fails for:

```
± |master ✓ | → go generate ./...
Writing `FakeThingDoer` to `counterfeiterreprofakes/fake_thing_doer.go`...
-: go: finding github.com/jfmyers9/counterfeiterrepro/counterfeiterreprofakes latest
go build github.com/jfmyers9/counterfeiterrepro/counterfeiterreprofakes: no Go files in
exit status 1
stuff.go:3: running "go": exit status 1
```

This is a common use case for repositories that do not check in their generated
fakes.
