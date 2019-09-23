package pkg

type Language uint8

const (
	Javascript Language = iota
	Rust
	Unknown
)

func getExtension(language Language) string {
	switch language {
	case Javascript:
		return "js"
	case Rust:
		return "rs"
	default:
		return ""
	}
}

func GetLanguageFromExtension(ext string) Language {
	switch ext {
	case "js":
		return Javascript
	case "rs":
		return Rust
	default:
		return Unknown
	}
}
