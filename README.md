# buildgo
Cross compile a .go file to a compiled Windows, Linux, and Darwin (Mac OS) executibles from your Mac or Linux computer automagically with the official golang docker container from hub.docker.com.

Requires Linux or Mac with Docker installed

A internet connection to download golang image from hub.docker.com

Example: 
"buildgo helloworld.go" will create 3 executibles:

helloworldlinuxamd64

helloworlddarwinamd64

helloworldwindowsamd64.exe
