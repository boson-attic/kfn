package pkg

type RuntimeManager interface {
	CheckDeps() bool
	Download()
}
