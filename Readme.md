# **Local File Server And Http File Proxy**

## Start :

There are three way for start program. Choose one of them.

1. With Source code
>go run main.go

By default program start **8888** port and **current folder**

you can change these setting with

> go run main.go -p="PORTNUMBER" -d="DIRECTORYROOTPATH"

2. With Complied version (for windows)
> ./hrassesment.exe

3. With Docker 

> docker run -it

## Usage :

After start program you can use ( http://localhost:8888 ) endpoint. In addition you can use http file such as given below ( http://localhost:8888?path=URL_THAT_YOU_WANT_TO_PROXY )

Local File Server Sample 
> http://localhost:8888

Video url Sample 
> http://localhost:8888/?path=https://sample-videos.com/video312/mp4/720/big_buck_bunny_720p_1mb.mp4

Http Text file url Sample
> http://localhost:8888/?path=https://sample-videos.com/text/Sample-text-file-1000kb.txt

## Tests :
>`go test -v`

result sample ;

![Result Sample](https://github.com/ysBayram/MiniFileServer/blob/master/testResult.PNG)

## Benchmark :

For Benchmark test use bombardier

First get bombardier with 

> `go get -u github.com/codesenberg/bombardier`

Then start 125 concurrent 100000 request with

> `bombardier -c 125 -n 100000 127.0.0.1:8888`

benchmark sample

![Benchmark Sample](https://github.com/ysBayram/MiniFileServer/blob/master/BenchResult.PNG)

### Test ,Benchmark And Development System Summary

- System Information report written at: 06/10/20
- OS Name	Microsoft Windows 10 Pro	
- Version	10.0.18362 Build 18362	
- OS Manufacturer	Microsoft Corporation	
- System Manufacturer	Hewlett-Packard	
- System Model	HP EliteBook Folio 9470m
- Processor	Intel(R) Core(TM) i7-3687U CPU @ 2.10GHz, 2601 Mhz, 2 Core(s), 4 Logical Processor(s)	
- BIOS Version/Date	Hewlett-Packard 68IBD Ver. F.47, 08/10/2013
- Time Zone	Turkey Standard Time	
- Installed Physical Memory (RAM)	8,00 GB	

## **Detail:**

This challenge is about building an API server serving files hosted on the local filesystem to the user. The project should serve files from a specific directory and subdirectories, and not anywhere outside. Implement an API endpoint for serving the files. The API will have an endpoint that'll receive the file path as a parameter, and return the contents to the user in an efficient manner. Think of the following issues while designing your backend: 
- how to specify configuration (location of the directory and other configuration you think is needed)
- how to serve huge files to users (>1GB) 
- how to scale up the number of users and support more concurrent requests

Implement the project in your language of choice, but include the whole project source including build tools, dependency manager, tests, and documentation. Your project will be evaluated based on readability, design, ease of use, and adherence to standards in addition to design and implementation correctness. 

If you manage to finish the project in time, as a stretch goal implement support for remote Http files: If the user sends an Http URL, fetch the URL and serve it to the user, acting as a simple Http proxy. 
