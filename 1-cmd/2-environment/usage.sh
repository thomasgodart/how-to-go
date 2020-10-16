#!/bin/bash

./environment --help
./environment -help
./environment --h
./environment -h

./environment

./environment --word=baz --numb=100 --bool=true
./environment --word baz --numb 100 --bool=true

./environment --word=test
./environment --word test
./environment -word=test
./environment -word test
./environment --word="test space"
./environment --word=test\ space
./environment --word="test \"quotation mark\""
./environment --word=test\ \"quotation\ mark\"

./environment -numb 1234 # decimal notation
./environment -numb 0664 # octal notation
./environment -numb 0x1234 # hexadecimal notation

./environment -bool # this bool will be "true"
./environment -bool=1
./environment -bool=0
./environment -bool=true
./environment -bool=false
./environment -bool=t
./environment -bool=f
./environment -bool=T
./environment -bool=F
./environment -bool=TRUE
./environment -bool=FALSE
./environment -bool=True
./environment -bool=False
