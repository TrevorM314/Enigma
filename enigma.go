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
var reflector = "YRUHQSLDPXNGOKMIEBFZCWVJAT";
// The index of the rotor key within rotors array
var firstRotor, secondRotor, thirdRotor int;
var rotorRotations[3] int;
var ringSettings[3] int;

func main() {
	setDefaults();
	message := ""
	for i := 65; i <=90; i++ {
		ch := string(i);
		encodedCh := encodeChar(ch);
		message = message + encodedCh;
	}
	fmt.Println(message);

	setDefaults();
	decoded := ""
	for i:=0; i<26; i++ {
		ch := string(message[i]);
		decodedCh := encodeChar(ch);
		decoded = decoded + decodedCh;
	}
	fmt.Println(decoded);
}

func encodeChar(c string) string {
	// Substitute character with plugboard match
	plugChar := plugboard[c];
	if plugChar == "" { 
		plugChar = c 
	}

	// Pass character through rotors
	rot1InAscii := int(plugChar[0]) - rotorRotations[0];
	rot1InAscii = fitAsciiToAlpha(rot1InAscii);
	rot1Out := rotorInToOut(0, string(rot1InAscii));

	rot2InAscii := int(rot1Out[0]) + rotorRotations[0] - rotorRotations[1];
	rot2InAscii = fitAsciiToAlpha(rot2InAscii);
	rot2Out := rotorInToOut(1, string(rot2InAscii));

	rot3InAscii := int(rot2Out[0]) + rotorRotations[1] - rotorRotations[2];
	rot3InAscii = fitAsciiToAlpha(rot3InAscii);
	rot3Out := rotorInToOut(2, string(rot3InAscii));

	// Pass through the reflector
	outIdx := ( int(rot3Out[0]) - 65 + rotorRotations[2] ) % 26;
	rot3InAscii = int(reflector[outIdx]) - rotorRotations[2];
	rot3InAscii = fitAsciiToAlpha(rot3InAscii);

	// Pass backward through rotors
	rot3Out = rotorOutToIn(2, string(rot3InAscii));

	rot2InAscii = int(rot3Out[0]) + rotorRotations[2] - rotorRotations[1];
	rot2InAscii = fitAsciiToAlpha(rot2InAscii);
	rot2Out = rotorOutToIn(1, string(rot2InAscii));

	rot1InAscii = int(rot2Out[0]) + rotorRotations[1] - rotorRotations[0];
	rot1InAscii = fitAsciiToAlpha(rot1InAscii);
	rot1Out = rotorOutToIn(0, string(rot1InAscii));

	// Pass backward through plugboard
	plugCharInAscii := int(rot1Out[0]) + rotorRotations[0];
	plugCharInAscii = fitAsciiToAlpha(plugCharInAscii);
	plugChar = plugboard[string(plugCharInAscii)];
	if plugChar == "" { plugChar = string(plugCharInAscii) }

	// Rotate rotors
	rotorRotations[0] ++;
	if rotorRotations[0] > 25 {
		rotorRotations[0] = 0;
		rotorRotations[1] ++;
	}
	if rotorRotations[1] > 25 {
		rotorRotations[1] = 0;
		rotorRotations[2] ++;
	}
	if rotorRotations[2] > 25 {
		rotorRotations[2] = 0;
	}

	return plugChar;
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
	rotorRotations = [3]int {2, 22, 1}

	/*
	Set rotor ringSettings
	integer 0-25, specifies how much the key should be shifted right (end chars moving to beginning).
	0 is the default setting and == the key found in rotors array
	*/
	ringSettings = [3]int {0, 0, 0}
}

func rotorInToOut(rotor int, in string) string {
	outIdx := int(in[0]) - 65;
	return string( rotors[rotor][outIdx] );
}

func rotorOutToIn(rotor int, out string) string {
	inASCII := strings.Index(rotors[rotor], out) + 65;
	return string(inASCII);
}

func fitAsciiToAlpha(value int) int {
	out := value
	for out < 65 {
		out += 26;
	}
	for out > 90 {
		out -= 26;
	}
	return out;
}