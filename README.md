Word Counter

This is a simple Go program that reads a text file and counts the frequency of each word in the file. It will output the most frequently occurring words in descending order, showing the top 5 words by default.

Features:

Reads a .txt file and counts word occurrences.

Outputs the top 5 most frequent words (or fewer if there aren't that many words).

Handles punctuation and converts words to lowercase for accurate counting.

Easy to configure the file path and top word limit.


File Structure:

wordCounter/

├── data/

│   ├── camus.txt

│   ├── test.txt

│   ├── test_data.txt

├── go.mod

├── main.go

data/: Contains the text files (camus.txt, test.txt, and test_data.txt) that are used for word counting.

go.mod: Go module configuration file.

main.go: The main Go program that performs the word counting.

Usage

The program reads the file specified in main.go (by default ./data/test_data.txt).

It will output the word frequencies to the console, listing the top 5 most frequent words.

You can modify the text file in the data/ folder and adjust the main.go file to specify a different file or top word limit if needed.

Example Output:

Top 5 most used words:

the: 23

a: 16

that: 16

is: 12

and: 12

License

This project is open source and available under the MIT License.
