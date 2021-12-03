package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	rand.Seed(42);
	secretMessage := encode("Hello world");
	fmt.Println(secretMessage);

	rand.Seed(42);
	message := encode(secretMessage);
	fmt.Println(message);
}

func encode(message string) string {
	message = strings.ToUpper(message)
	encoded := "";
	for i := 0; i < len(message); i++ {
		plugboard := scramble();
		encodedCh := plugboard[string(message[i])];
		if encodedCh == "" { encodedCh = string(message[i]) }
		encoded = encoded + encodedCh;
	}
	return encoded;
}

func scramble() map[string]string {
	plugboard := make( map[string]string );
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	for len(alphabet) > 0 {
		pair := rand.Int() % ( len(alphabet)-1 ) + 1;
		plugboard[string(alphabet[0])] = string(alphabet[pair]);
		plugboard[string(alphabet[pair])] = string(alphabet[0]);
		alphabet = alphabet[1:pair] + alphabet[pair+1:];
	}
	return plugboard
}