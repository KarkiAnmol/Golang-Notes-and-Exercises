// -------------------------------NOTE-----------
// Strings and Runes and Bytes
var s string = "Hello there"
var b byte = s[6] // valid 

var s string = "Hello there" //here code point in UTF-8 can be anywhere from 1 to 8 bytes long
// slicing to make substrings
var s2 string = s[4:7] // s2 = "o t"
var s3 string = s[:5] // s3= "Hello"
var s4 string = s[6:] // s4= "there"

//when using emojis, code point in UTF-8 can be several bytes long
var s string = "Hello 🐍" // len(s)= 4 bytes long
var s3 string = s[:5] // "Hello"	
var s2 string = s[4:7] // "o ?" --It's because we only copied the first byte of the emoji's code point which is invalid

//Strings are immutable

//rune to String
var a rune = 'x'
var s string = string(a)

//byte to String
var b byte = 'y'
var s2 string = string(b)

//Converting strings to slices
var s string = "Hello 🐍" 
var bs []byte = []byte(s)
var rs []rune = []rune(s)
fmt.Println(bs)
fmt.Println(rs)
// When you run this code, you see:
[72 101 108 108 111 44 32 240 159 140 158] //String converted to UTF-8 bytes
[72 101 108 108 111 44 32 127774] //String converted to runes

var s string = "🌞"
var bs []byte = []byte(s)
var rs []rune = []rune(s)
fmt.Println(len(s))
fmt.Println(bs)
fmt.Println(rs)
// When you run this code, you see:
  4 //length of emoji is 4 bytes unlike characters which are of 1 byte
[240 159 140 158]
[127774]

//********** UTF-8 *********
// --> UTF-32 wastes too much space (11 digits always 0)
// -->UTF-16 also wasteful
// -->UTF-8 uses a single byte(1) to represent the Unicode characters
// 		whose values are below 128 (which includes all of the letters, numbers, and punctuation commonly used in English), 
//      but expands to a maximum of four bytes(4) to represent Unicode code points with larger values
// --> pros: I) no need to worry about little endian vs big endian
//			    II) It also allows you to look at any byte in a sequence and tell if you are at the start of a
//               UTF-8 sequence, or somewhere in the middle
// --> cons: I) you cannot randomly access a string encoded with UTF-8.
//		     	II) While you can detect if you are in the middle of a character, you can’t tell how many
//				characters in you are.You need to start at the beginning of the string and count.

