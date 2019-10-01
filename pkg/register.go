package pkg

import (
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/languages/js"
)

func init() {
	languages.RegisterLanguageManager(languages.Javascript, js.NewJsLanguageManger())
}
