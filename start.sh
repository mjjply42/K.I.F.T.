#!/bin/sh

# COLORS
Black=$'\e[1;30m'
Blue=$'\e[1;34m'
Green=$'\e[1;32m'
Cyan=$'\e[1;36m'
Red=$'\e[1;31m'
Purple=$'\e[1;35m'
Brown=$'\e[1;33m'
Gray=$'\e[1;37m'
Dark_Gray=$'\e[1;30m'
Light_Blue=$'\e[1;34m'
Light_Green=$'\e[1;32m'
Light_Cyan=$'\e[1;36m'
Light_Red=$'\e[1;31m'
Light_Purple=$'\e[1;35m'
Yellow=$'\e[1;33m'
White=$'e[1;37m'
End=$'\e[0m'

# Regular Variables
CHECK=$(which brew)
POCKETSPHINX=$(brew info pocketsphinx | grep '/usr')

if [ -z "$CHECK" ]
then
    echo "${Red} Homebrew does not exist${End}....${Yellow}checking dependencies${End}"  
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
else
      echo "${Green} Homebrew exists${End}....${Yellow}checking dependencies${End}"
fi

if [ -z "$POCKETSPHINX" ]
then
      cd pocketsphinx
      make setup
      make re
      cd ..
else
      cd pocketsphinx
      make re
      cd ..
fi
echo ${Blue} Starting Program...${End}
go run main.go