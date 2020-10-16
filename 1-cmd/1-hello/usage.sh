#!/bin/bash

# "--" or "-" is equivalent:

./hello --help
./hello -help

# "help" or "h" is equivalent:

./hello --h
./hello -h

# you get help automatically if your command is wrong:

./hello -wrong -command

# this works:

./hello

# "=" is optional, except for booleans

./hello --word=baz --numb=100 --bool=true
./hello --word baz --numb 100 --bool=true

./hello --word=test
./hello --word test
./hello -word=test
./hello -word test
./hello --word="test space"
./hello --word=test\ space
./hello --word="test \"quotation mark\""
./hello --word=test\ \"quotation\ mark\"

./hello -numb 1234 # decimal notation
./hello -numb 0664 # octal notation
./hello -numb 0x1234 # hexadecimal notation

# booleans only always need the "="

./hello -bool # this bool will be "true"
./hello -bool=1
./hello -bool=0
./hello -bool=true
./hello -bool=false
./hello -bool=t
./hello -bool=f
./hello -bool=T
./hello -bool=F
./hello -bool=TRUE
./hello -bool=FALSE
./hello -bool=True
./hello -bool=False
