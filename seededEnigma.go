package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"bufio"
	"os"
)

var (
	rng = rand.New(rand.NewSource(int64(42)))
)

func main() {
	//Read in Seed
	var input string;
	fmt.Println("Select a seed (int)");
	fmt.Scanln(&input);
	seed, err := strconv.ParseInt(input, 10, 64);
	for err != nil {
		fmt.Println("Select a seed (int)");
		fmt.Scanln(&seed);
		seed, err = strconv.ParseInt(input, 10, 64);
	}
	fmt.Println(seed)

	rng.Seed(seed);

	var message string;
	for len(message) < 1 {
		fmt.Println("Enter your message to be encoded");
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message = scanner.Text()
		message = strings.ToUpper(message);
	}

	encoded := encode(message);
	fmt.Println(encoded);
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
		pair := rng.Int() % ( len(alphabet)-1 ) + 1;
		plugboard[string(alphabet[0])] = string(alphabet[pair]);
		plugboard[string(alphabet[pair])] = string(alphabet[0]);
		alphabet = alphabet[1:pair] + alphabet[pair+1:];
	}
	return plugboard
}