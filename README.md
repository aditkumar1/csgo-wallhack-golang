# CSGO Wallhack Golang
External glow ESP for CS:GO
<br>

![ESP Screenshot](https://raw.githubusercontent.com/aditkumar1/csgo-wallhack-golang/main/screenshots/20220207015205_1.jpg)

## Prerequisites
<li>Install <a href="https://go.dev/doc/install">Golang</a></li>
<li>Admin privileges on local machine</li>

## Getting Started:

Step 1: Getting this repo in your local machine.
```
Clone this repo  "git clone https://github.com/aditkumar1/csgo-wallhack-golang.git". Alternatively download ZIP.
```

Step 2: Running the cheat
<br>
<br>
Note-1) Launch CSGO application before running below command
<br>
Note-2) Below command needs to run with admin privileges
<br>

```
go run hack.go
```

## Tested on
<li>GO 1.14</li>
<li>Windows 10</li>

## Known Errors
```
 Encountered error - runtime error: invalid memory address or nil pointer dereference , ensure CS GO application running before starting this application.
```
This is self explanatory and indicates that CSGO process needs to start before running this cheat

## TO DOs
<li>Feature to update offsets on application start</li>
<li>Better application logging</li>

## Thanks to: 
<li><a href = "https://github.com/Xustyx/goxymemory"> Goxymemory</a> for providing wrapper around win32 process read memory and write memory
<br/></li>

## Enjoy the game

![ESP Screenshot](https://raw.githubusercontent.com/aditkumar1/csgo-wallhack-golang/main/screenshots/20220207015211_1.jpg)