package shared

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func CountKeyStrokes() int {
	var counter int
	var typedCharacters []rune
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Presione ENTER dos veces para salir")
	fmt.Println("Ingrese una descripción: ")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeySpace {
			fmt.Print(" ")
			typedCharacters = append(typedCharacters, ' ')
		} else if key == keyboard.KeyBackspace && len(typedCharacters) > 0 {
			fmt.Print("\b \b")
			typedCharacters = typedCharacters[:len(typedCharacters)-1]
		} else if char != 0 {
			fmt.Printf("%c", char)
			typedCharacters = append(typedCharacters, char)
		}
		counter++
		if key == keyboard.KeyEnter {
			break
		}
	}
	return counter
}
