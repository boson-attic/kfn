package pkg

import (
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/languages/js"
	"github.com/slinkydeveloper/kfn/pkg/languages/rust"
)

func init() {
	languages.RegisterLanguageManager(languages.Javascript, js.NewJsLanguageManger())
	languages.RegisterLanguageManager(languages.Rust, rust.NewRustLanguageManger())
}
