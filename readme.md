# PromoGo

This project has two modes: "generate" and "listen". In "generate" mode, a file of promotions will be generated, and in "listen" mode, the HTTP server will listen for requests and serve promotions based on the ID given.

## Usage

To build the binary, run the following command

`go build`

This command will create a binary file named `problem-1`

### Generate promotions

To generate a promotions CSV file, run the following command:

`./problem-1 generate --num=<number of promotions> --file=<filename.csv>`

This command will generate a file with a given number of promotions, and save it to a specified filename.

### Start http server

To listen to the HTTP server, run the following command:

`./problem-1 listen --port=<port number> --file=<filename.csv>`

This command will start the server on the specified port, and serve promotions based on the ID given.

## Code

The "main" function reads in the command line arguments and uses a switch statement to determine which mode to run in. If the "generate" flag is set, the "GeneratePromotionsCsvFile" function is called to generate the promotions file. If the "listen" flag is set, the "listen" function is called to start the HTTP server.

The "listen" function initializes the storage with the CSV file and sets a time limit for storage. The "RunHttpServer" function is then called to start the HTTP server on the specified port. Finally, the function closes the storage when the server is stopped.

To extract the required line from the promotions file quickly, the project first calculates all the offsets of the lines in the file and stores them in memory. An offset is the byte position of the start of a line in the file. By storing all the offsets in memory, the project avoids the need to search the file every time it needs to extract a specific line.

When a request is made for a specific promotion based on the ID given, the project uses the calculated offsets to quickly extract the corresponding line from the file. It does this by calculating the byte position of the start of the line based on the ID, and then reading the corresponding number of bytes from the file.

Overall, this approach allows the project to extract specific lines from the file quickly and efficiently, without the need to search the file every time a request is made.