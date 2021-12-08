package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
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
var rotorSelections[3] int;

var rotorRotations[3] int;
var ringSettings[3] int;

func main() {
	setDefaults();
	getSettings();

	var message string;
	for len(message) < 1 {
		fmt.Println("Enter your message to be encoded");
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message = scanner.Text()
		message = strings.ToUpper(message);
	}

	encoded := ""
	for i := 0; i < len(message); i++ {
		ch := string(message[i]);
		encodedCh := encodeChar(ch);
		encoded = encoded + encodedCh;
	}
	fmt.Println(encoded);
}

func encodeChar(c string) string {
	// Don't encode non alphabetic characters
	if int(c[0]) < 65 || int(c[0]) > 90 {
		return c;
	}

	// Substitute character with plugboard match
	plugChar := plugboard[c];
	if plugChar == "" { 
		plugChar = c 
	}

	// Pass character through rotors
	rot1InAscii := int(plugChar[0]) - rotorRotations[0];
	rot1InAscii = fitAsciiToAlpha(rot1InAscii);
	rot1Out := rotorInToOut(rotorSelections[0], string(rot1InAscii));

	rot2InAscii := int(rot1Out[0]) + rotorRotations[0] - rotorRotations[1];
	rot2InAscii = fitAsciiToAlpha(rot2InAscii);
	rot2Out := rotorInToOut(rotorSelections[1], string(rot2InAscii));

	rot3InAscii := int(rot2Out[0]) + rotorRotations[1] - rotorRotations[2];
	rot3InAscii = fitAsciiToAlpha(rot3InAscii);
	rot3Out := rotorInToOut(rotorSelections[2], string(rot3InAscii));

	// Pass through the reflector
	outIdx := ( int(rot3Out[0]) - 65 + rotorRotations[2] ) % 26;
	rot3InAscii = int(reflector[outIdx]) - rotorRotations[2];
	rot3InAscii = fitAsciiToAlpha(rot3InAscii);

	// Pass backward through rotors
	rot3Out = rotorOutToIn(rotorSelections[2], string(rot3InAscii));

	rot2InAscii = int(rot3Out[0]) + rotorRotations[2] - rotorRotations[1];
	rot2InAscii = fitAsciiToAlpha(rot2InAscii);
	rot2Out = rotorOutToIn(rotorSelections[1], string(rot2InAscii));

	rot1InAscii = int(rot2Out[0]) + rotorRotations[1] - rotorRotations[0];
	rot1InAscii = fitAsciiToAlpha(rot1InAscii);
	rot1Out = rotorOutToIn(rotorSelections[0], string(rot1InAscii));

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
	rotorSelections = [3]int {0, 1, 2}

	/*
	Set starting Rotor rotations
	integer 0-25
	*/
	rotorRotations = [3]int {2, 25, 0}

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

func getSettings() {
	var input string;

	// Select Rotors & Rotation
	for i := 0; i < 3; i++ {
		//Select Rotor
		selectionValid := 0;
		for selectionValid != 1 {
			input = ""
			fmt.Printf("Select Rotor #%d (0-7) [default %d]: \n", i, i)
			fmt.Scanln(&input);
			var selection int;
			var err error;
			if input == "" { // Use default
				selection = i
			} else {
				selection, err = strconv.Atoi(input);
			}

			if err != nil {
				fmt.Println("Unable to parse input. Please only enter a number 0-7");
			} else if selection < 0 || selection > 7 {
				fmt.Println("Selection out of bounds");
			} else {
				for j := 0; j<i; j++ {
					if rotorSelections[j] == selection {
						selectionValid = -1
					}
				}
				if selectionValid == -1 {
					fmt.Printf("Rotor %s alrady in use\n", strconv.Itoa(selection));
					selectionValid = 0;
				} else { // VALID
					rotorSelections[i] = selection;
					selectionValid = 1;
				}
			}
		}

		//Select Rotation
		selectionValid = 0;
		for selectionValid != 1 {
			input = ""
			fmt.Println("Select starting rotation (0-25) [default 0]: ")
			fmt.Scanln(&input);
			if input == "" { // Use default
				selectionValid = 1;
				continue;
			}
			selection, err := strconv.Atoi(input);
			if err != nil {
				fmt.Println("Unable to parse input. Please only enter a number 0-25");
			} else if selection < 0 || selection > 25 {
				fmt.Println("Selection out of bounds");
			} else {
				rotorRotations[i] = selection;
				selectionValid = 1;
			}
		}
	}

	// Plugboard settings 
	fmt.Println("Enter plugboard settings separated by space (ie 'AZ BR TN QP') [default AB CD ... ST]")
	plugboardValid := 0

	for plugboardValid != 1 {
		input = ""
		// Use bufio isntead of scanf to allow spaces
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		input = strings.ToUpper(input);
		if input == "" { // use default
			plugboardValid = 1;
			continue;
		}

		var tempPlugs = make( map[string]string );
		//Modes:
		EXPECTFIRST := 0;
		EXPECTSECOND := 1;
		EXPECTSPACE := 2;
		FAIL := 3;
		mode := EXPECTFIRST
		firstChar := ""
		for ch := 0; ch < len(input); ch ++ {
			if mode == EXPECTFIRST {
				if string(input[ch]) == " " {
					continue;
				} else if input[ch] < 65 || input[ch] > 90 {
					fmt.Printf("unexpected char at %d: %s\n", ch, string(input[ch]));
					mode = FAIL;
					break;
				} else {
					firstChar = string(input[ch])
					mode = EXPECTSECOND;
				}
			} else if mode == EXPECTSECOND {
				if input[ch] < 65 || input[ch] > 90 {
					fmt.Printf("unexpected char at %d: %s\n", ch, string(input[ch]));
					mode = FAIL;
					break;
				} else {
					secondChar := string(input[ch]);
					if tempPlugs[firstChar] != "" {
						fmt.Printf("Character '%s' used twice\n", firstChar);
						mode = FAIL;
						break;
					} else if tempPlugs[secondChar] != "" {
						fmt.Printf("Character '%s' used twice\n", secondChar);
						mode = FAIL;
						break;
					} else { // Valid pair
						tempPlugs[firstChar] = secondChar;
						tempPlugs[secondChar] = firstChar;
						mode = EXPECTSPACE;
					}
				}
			} else if mode == EXPECTSPACE {
				if string(input[ch]) != " " {
					fmt.Printf("Error, expected space but read '%s' at %d\n", string(input[ch]), ch)
					mode = FAIL;
					break;
				} else {
					mode = EXPECTFIRST;
				}
			}
		}
		// Deep copy tempPlugs to plugboard
		if mode == EXPECTSPACE || mode == EXPECTFIRST {
			for k,v := range tempPlugs {
				plugboard[k] = v
			}
			plugboardValid = 1;
		} else {
			plugboardValid = 0;
		}
	}
}