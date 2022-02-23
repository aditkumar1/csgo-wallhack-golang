# CSGO Wallhack Golang
External glow ESP for CS:GO
<br>

![ESP Screenshot](https://raw.githubusercontent.com/aditkumar1/csgo-wallhack-golang/main/screenshots/20220207015205_1.jpg)

## Prerequisites
<li>Install <a href="https://go.dev/doc/install">Golang</a></li>
<li>Admin privileges on local machine</li>

## Getting Started:

Step 1: Getting this repo in your local machine. <br><br>
Clone this repo -
```
git clone https://github.com/aditkumar1/csgo-wallhack-golang.git
```
<br/>
Alternatively download ZIP by clicking on Code -> Download Zip.
<br/><br/>
Step 2: Running the cheat
<br>
<br>
<br>
Note-1) Launch CSGO application before running below command. VAC ban risk is always there. On safe side, start your CS GO game first with -insecure flag
<br>
<br>
Note-2) Below command needs to run with admin privileges
<br>
<br>
Note-3) Make sure you are inside root directory of this cheat by running.

```
cd csgo-wallhack-golang
```


Run it

```
go run hack.go
```
<br>

Alternatively, download and run using go get (preferred if you can run application in your GOPATH) - 
<br/>

```
go get -v github.com/aditkumar1/csgo-wallhack-golang
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
<li>Improve performance by invoking garbace collector efficiently.</li>
<li>Improve performance by doing less type conversions during runtime.
<li>Better application logging.</li>
<li> Implement better evasive techniques to bypass VAC.

## Opening an issue
<li>Provide any specific error you are seeing. Screenshot will help</li>
<li>Provide output of following command -

```
go env
```

<li>Please dont use any swearing language and be patient. We will work this out together.</li>
<li> As always, any contributions will be helpful.

## Thanks to: 
<li><a href = "https://github.com/Xustyx/goxymemory"> Goxymemory</a> for providing wrapper around win32 process read memory and write memory
<br/></li>

## Enjoy the game

![ESP Screenshot](https://raw.githubusercontent.com/aditkumar1/csgo-wallhack-golang/main/screenshots/20220207015211_1.jpg)