package main

import "fmt"

const helloPrefixEn = "Hello, "
const helloPrefixDe = "Hallo, "
const worldEn = "World! "
const worldDe = "Welt!"

func main() {
	fmt.Print(Hello("Andrei!", "DE"))
}

func Hello(name, language string) string {
	if name == "" && language == "" {
		return "Hello, World!"
	}
	if name != "" {
		return getPrefixInLanguage(language) + name + "!"

	} else {
		return getPrefixInLanguage(language) + getWorldOnLanguage(language)
	}
}

func getWorldOnLanguage(language string) string {
	switch language {
	case "DE":
		return worldDe
	case "EN":
		return worldEn
	}
	return ""
}

func getPrefixInLanguage(language string) string {
	switch language {
	case "DE":
		return helloPrefixDe
	case "EN":
		return helloPrefixEn
	}
	return ""
}
