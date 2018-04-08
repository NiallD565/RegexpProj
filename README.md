# RegexpProj

Requirements :
  1. Convert a infix expression to a postfix expression
  2. Build a non-deterministic finite automaton (NFA) from a regular expression.
  3. Compare two strings and see if they match using a regular expressions
  
# Description

The aim of this project is to create a program in the Go programming language that can build  non-determiniatic 
finite automaton (NFA) from a regular expression, and can use a the NFA to check if the reular expression matches 
any given string of text. It can not use the regexp package from the go standard library or any other external library.
A regular expression is a string containing a series of characters that have a special meaning, in this program I have 
used the ".","*","|" to denote concatinate, kleene star and or.

# Planning and Development

I began my project by researching regular expression engines and reviewing the lecture notes. that gave me the initial 
idea how the project will work. My initial idea was to encapsulate the prject but after a small amount of research I 
found out Go is not object orientated. 

Using the video tutorials provided I created similar classes and methods to the ones made in the videos I then adapted them
into one class and created a user input method so the user can enter a infix expression as described in the project brief 
which will then be converted to postfix

# Research 

To begin i looked at Golangs Regexp documentation. From that in understood that a regular expression was a series of characters
expressing a string or pattern to be searched for within a larger piece of text. From the video lectures explaining thompsons
constructor and the shunting yard algorithm I grasped that what the program was suppose to do

# Running the program

This program is written in the Go programming language.
If you don't have Go installed you can follow the link https://golang.org/doc/install

To clone this repository enter the following command into your command line 
git clone https://github.com/NiallD565/RegexpProj

To run the program you will need to navigate to where the file was saved and run this command
go run .\Regexp.go

# Design
This program was developed using the Go programming Language, Visual Studio Code and Git

Technologies Used
- GoLang
- Visual Studio Code
- Git

# Resources 
https://web.microsoftstream.com/video/84d5b1a6-716a-49d6-abaf-2a9c2790cd4b
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
