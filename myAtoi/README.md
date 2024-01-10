## myAtoi
This problem requires to implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).
The solution provided in this repo is implemented in Go.

### The algorithm for myAtoi(string s) is as follows:

1) Read in and ignore any leading whitespace.
2) Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
3) Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
4) If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
5) Return the integer as the final result.

### Note:
Only the space character ' ' is considered a whitespace character.\
__Do not ignore__ any characters other than the leading whitespace or the rest of the string after the digits.

### Some possible testcases:
" -0012a42"\
"00000-42a1234"\
"  0000000000012345678"\
"  +  413"\
"+-12"\
"   -42"\
"-000000000000000000000000000000000000000000000000001"\
"20000000000000000000"\
"words and 987"\
"-5-"

### Details of the submission:
Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).\
Memory Usage: 2.2 MB, less than 99.86% of Go online submissions for String to Integer (atoi).\
Sumbitted on 10.01.2024
