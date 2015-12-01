wShare
========

wShare was built to make file sharing within intranet easy.

##Build

1. `go get github.com/thewhitetulip/wshare`
2. `go build`


##Working

Call the binary in the following way:

`./wshare -f <filename> -t <number>`

eg:   `./wshare -f classical.pdf`

where `filename` is the name of the file and `number` is the number of 
times the file is expected to be downloaded.

When you run the command as above, it'll give a URL like
 
>download link: http://127.0.0.1:8080/share/classical.pdf

>Running server on port 8080

this link can be used to download the file over any browser or a 
download manager if you are using *nix and have wget then do a 

`wget http://127.0.0.1:8080/share/classical.pdf`

##TODO
1. Ability to share folders
2. Short URLs (will involve routing)


##Contributing:

Pull requests are welcome. If you are a first time open source contributor, 
don't hesitate to open an issue, and let's discuss, I don't bite!

Please run gofmt on all go files you change or submit, see 
`github.com/thewhitetulip/gofmtall` for a shell script which runs gofmt on each go file in 
the directory

LICENSE: MIT