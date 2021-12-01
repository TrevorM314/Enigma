package main

import (
	"fmt"
	"strings"
)

/*
General Notes
Each time a key is pressed, the rotation of firstRotor ++;
if firstRotor > 25, firstRotor = 0; secondRotor ++ (etc. for third);
By typing encrypted message into enigma with same start settings, decrypted message revealed
*/

// Hashmap where each character is the key and it's paired character is the value
var plugboard = make( map[string]string );
// Rotor wiring source: https://en.wikipedia.org/wiki/Enigma_rotor_details#Rotor_wiring_tables
var rotors = [8]string {
	"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"VZBRGITYUPSDNHLXAWMJQOFECK",
	"JPGVOUMFYQBENHZRDKASXLICTW",
	"NZJHGRCXMYSWBOUFAIVLPEKQDT",
	"FKQHTLXOCBJSPDZRAMEWNIUYGV",
}
// The index of the rotor key within rotors array
var firstRotor, secondRotor, thirdRotor int;

func main() {
	setDefaults();
	fmt.Println(encodeChar("A"));
	fmt.Println(encodeChar("Z"));
}

func encodeChar(c string) string {
	// Substitute character with plugboard match
	plugChar := plugboard[c];
	if plugChar == "" { plugChar = c }

	// Pass character through rotors

	// Pass backward through rotors

	return "";
}

func setDefaults() {
	/*
	Set Plugboard
	For A-T, A<-->B, C<-->D, etc
	*/
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			plugboard[string(65 + i)] = string(65 + i + 1);
		} else {
			plugboard[string(65 + i)] = string(65 + i - 1);
		}
	}
	//Note! If a letter is not found in the plugboard, it stays as self.

	/*
	Set Rotors being used
	*/
	firstRotor = 0;
	secondRotor = 1;
	thirdRotor = 2;

	/*
	Set starting Rotor rotations
	integer 0-25
	*/

	/*
	Set rotor ringSettings
	integer 0-25, specifies how much the key should be shifted right (end chars moving to beginning).
	0 is the default setting and == the key found in rotors array
	*/
}

func rotorInToOut(rotor int, in string) string {
	outIdx := int(in[0]) - 65;
	return string( rotors[rotor][outIdx] );
}

func rotorOutToIn(rotor int, out string) string {
	inASCII := strings.Index(rotors[rotor], out) + 65;
	return string(inASCII);
}