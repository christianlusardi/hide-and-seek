package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/auyer/steganography"
)

var pictureInputFile string
var pictureOutputFile string
var messageInputFile string
var messageOutputFile string
var decode bool
var encode bool
var noBanner bool
var help bool

// init creates the necessary flags to run program from the command line
func init() {

	flag.BoolVar(&decode, "d", false, "Flag to decode a message from a given PNG file")
	flag.BoolVar(&encode, "e", false, "Flag to encode a message to a given PNG file")

	flag.StringVar(&pictureInputFile, "pi", "", "Path to the the input image")
	flag.StringVar(&pictureOutputFile, "po", "encoded.png", "Path to the the output image")

	flag.StringVar(&messageInputFile, "mi", "", "Path to the message input file")
	flag.StringVar(&messageOutputFile, "mo", "", "Path to the message output file")

	flag.BoolVar(&noBanner, "nobanner", false, "Flat to disable banner printing")

	flag.BoolVar(&help, "help", false, "Help")

	flag.Parse()
}

func main() {

	if !noBanner {
		//print banner
		banner()
	}

	//encode
	if encode {

		encodeImg()

	} else if decode {

		decodeImg()

	} else if help {

		helpPrint()

	} else {
		fmt.Println("Error operation not supported")
	}

}

func banner() {
	b, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func encodeImg() {
	message, err := ioutil.ReadFile(messageInputFile) // Read the message from the message file (alternative to os.Open )
	if err != nil {
		print("Error reading from file!!!")
		fmt.Println(err.Error())
		return
	}

	inFile, err := os.Open(pictureInputFile) // Opens input file provided in the flags
	if err != nil {
		log.Fatalf("Error opening file %s: %v", pictureInputFile, err)
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile) // Reads binary data from picture file
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	encodedImg := new(bytes.Buffer)
	err = steganography.Encode(encodedImg, img, message) // Calls library and Encodes the message into a new buffer
	if err != nil {
		log.Fatalf("Error encoding message into file  %v", err)
	}

	outFile, err := os.Create(pictureOutputFile) // Creates file to write the message into
	if err != nil {
		log.Fatalf("Error creating file %s: %v", pictureOutputFile, err)
	}
	bufio.NewWriter(outFile).Write(encodedImg.Bytes()) // writes file to disk

	fmt.Println("Task successfully finished!")
}

func decodeImg() {
	inFile, err := os.Open(pictureInputFile) // Opens input file provided in the flags
	if err != nil {
		log.Fatalf("Error opening file %s: %v", pictureInputFile, err)
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal("error decoding file", img)
	}

	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // Uses the library to check the message size

	msg := steganography.Decode(sizeOfMessage, img) // Read the message from the picture file

	// if the user specifies a location to write the message to...
	if messageOutputFile != "" {
		err := ioutil.WriteFile(messageOutputFile, msg, 0644) // write the message to the given output file

		if err != nil {
			fmt.Println("There was an error writing to file: ", messageOutputFile)
		}
	} else { // otherwise, print the message to STDOUT
		for i := range msg {
			fmt.Printf("%c", msg[i])
		}
	}

}

func helpPrint() {
	fmt.Println("How to use this script:")
	fmt.Println("-i: the input image to encode in / decode from")
	fmt.Println()
	fmt.Println("-e: take a message and encodes it into a specified location")
	fmt.Println("-mi: input message to for the encoding option 			(ENCODING ONLY)")
	fmt.Println("-o: where you would like to store the encodeded image		(ENCODING ONLY)")
	fmt.Println("\t+ EX: ./main -e -i plain.png -mi message.txt  -o secret.png")
	fmt.Println()
	fmt.Println("-d: take a picture and decodes the message from it")
	fmt.Println("-mo: output message. Lempty for STDIO			(DECODING ONLY)")
	fmt.Println("\t+ EX: ./stego -d -i secret.png -mo secret.txt")
}
