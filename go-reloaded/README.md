### Go-reloaded

The Program recieves unformatted text from a file process it and give out a clean and well arrange text suitable for usage.

### HOW IT WORK

- The Go-reloaded project is a full mini text editor, It work base on instance of commands like (up), (low), (cap),(hex) etc. And other special functions.

- Quote were explicitly handle by the fixquote function, wrapping either double quote or single quote round a word or sentence correctly, in a case where quotes are missing their pairs the program accurately detect which is an opening or closing quote.

- Every instance of a command like (low, 4) or (cap, 7) etc. should effect to the previous word accord to the number and the command, it handle cases where the commands were place at the very beginning of the sentance or test cases.

- It fix punctuation to previous word and also group punctuation for more readability.

- Commands like (hex) convert the Previous word from base 16 to base 10 while command like (bin) convert the Previous word from base 2 to base 10.

### What I Learnt From The Program

1. I learnt how to convert between bases.

2. Basic knowledge about (regexp package) in Golang.

3. String malnipulation.

4. Different way of reading and writing file, even outputing on a newline using bufio.

5. Learnt file structure.

6. How to initialize go mod. 

7. Learnt basics of error handling 