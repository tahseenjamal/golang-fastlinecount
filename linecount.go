// +build

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//Get file name from the command line argument
	filename := os.Args[1]

	//Create file handle for the file, whose name has been passed through command line argument
	filehandle, err := os.Open(filename)

	//automatically close file handle once main() exits
	defer filehandle.Close()

	//Initialize linecount, this would be incremented to count lines
	linecount := 0

	//Buffer size for reading chunks of data block from the file. To be optimal, depending on your RAM size. This is default is good enough
	buffer_size := 1024 * 1024

	//data buffer created, this would be filled with block data. Create is once and reuse it in code
	data_buffer := make([]byte, buffer_size)

	//If no error in reading the file then enter in this IF condition
	if err == nil {

		//This is endless loop and would be broken out by end of file
		for {

			//read data from file in block
			size_received, err := filehandle.Read(data_buffer)

			//trim data buffer to the size of received data to avoid trailing 0
			data_buffer = data_buffer[0:size_received]

			//If error then it means nothing more to read, end of file reached.
			if err != nil {

				/*if size_received > 0 && strings.Count(string(data_buffer), "\n") > 0 {
					linecount++
				}
				*/
				break
			}

			//increment linecount by number of newline characters found
			linecount += strings.Count(string(data_buffer), "\n")

		}
	}

	fmt.Println(linecount)
}
