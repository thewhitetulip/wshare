# wShare
========

wShare was built to make file sharing within intranet easy.

## Build

1. `go get github.com/thewhitetulip/wshare`
2. `go build`


## Working

Call the binary in the following way:

`./wshare -f <filename> -t <number>`

eg:   `./wshare -f classical.pdf`

where `filename` is the name of the file and `number` is the number of 
times the file is expected to be downloaded.

If you want to share a folder, then provide the folder name after the `-f` flag, the folder will be
compressed and that zip file will be shared.

When you run the command as above, it'll give a URL like
 
>download link: http://127.0.0.1:8080/share/classical.pdf

>Running server on port 8080

this link can be used to download the file over any browser or a 
download manager if you are using *nix and have wget then do a 

`wget http://127.0.0.1:8080/share/classical.pdf`

## TODO
1. easy way to share with mobile devices which don't support zip

## Contributing:

Pull requests are welcome. If you are a first time open source contributor, 
don't hesitate to open an issue, and let's discuss, I don't bite!


wShare is intended to used for sharing files within the network, there are two ways you can share files, a single file 
mode and a multiple file mode, above mentioned is the example for single file mode, if you have to share multiple files then
use the directory mode and share the directory, the application will zip it and allow it to be downloaded, either that or compress
the respective files and share the one zip file, the choice is yours, the application is designed to suit to the user's preferecnes
and not the other way round!
 
LICENSE: MIT
