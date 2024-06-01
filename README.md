# Rot14 Encryption and Decryption

This program encrypts or decrypts a given input string using the Rot14 encryption technique.

## Usage

### Encrypting a String

To encrypt a string using Rot14 encryption, run the program with the `--to-rot14` flag followed by the input string enclosed in double quotes.

Example:
```
go run . --to-rot14 "input string"
```

### Decrypting a String

To decrypt a string encrypted with Rot14 encryption, run the program with the `--from-rot14` flag followed by the input string enclosed in double quotes.

Example:
```
go run . --from-rot14 "input string"
```

## Running the Program

Ensure you have Go installed on your machine. Clone this repository and navigate to its directory in your terminal. Then, execute the `go run .` command followed by the appropriate flag and input string.

## Note

- This program only accepts one argument as input.
- If excess or insufficient arguments are provided, the program will exit with an error message.
- Ensure that the input string is enclosed in double quotes.
- Flags must be either `--to-rot14` for encryption or `--from-rot14` for decryption.

## Example

**Encrypting a string:**
```
go run . --to-rot14 "Hello, World!"
```
Output:
```
Vszzc, Kcfzr!
```

**Decrypting a string:**
```
go run . --from-rot14 "Vszzc, Kcfzr!"
```
Output:
```
Hello, World!
```
## Requirements
``rot-14`` is a stand-alone program written on go. To run it, you will need to have Golang insalled in your computer

## Contributing to to-roman
If you encounter any issues or have suggestions for improvements, feel free to contribute to the project. You can submit bug reports or pull requests on the GitHub repository.
## License
This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License). Feel free to use, modify, and distribute the code for personal or commercial purposes.
## Author
David Jesse Odhiambo

Apprentice Software Developer, Zone01 Kisumu