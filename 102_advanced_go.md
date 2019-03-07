
# Golang programming course

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Golang programming course](#golang-programming-course)
    - [Why Golang?](#why-golang)
    - [Golang in short](#golang-in-short)
        - [Code](#code)
        - [Packages](#packages)
        - [Concurrency](#concurrency)
    - [Install Golang](#install-golang)
    - [Learning go](#learning-go)
        - [These are some useful introductions to Go:](#these-are-some-useful-introductions-to-go)
        - [Tips when learning to code](#tips-when-learning-to-code)
        - [Additional resources](#additional-resources)
        - [Nice excercises](#nice-excercises)
        - [Interesting Go talks](#interesting-go-talks)
        - [Use the power](#use-the-power)
        - [Testing](#testing)
        - [Reference](#reference)
    - [Must read stuff](#must-read-stuff)
    - [Protocols](#protocols)
        - [HTTP](#http)
        - [Websocket](#websocket)
    - [Basics](#basics)
    - [Go and Docker](#go-and-docker)
    - [Dep](#dep)

<!-- markdown-toc end -->


## Why Golang?
Go was created by Google out of need. Why don't we let the creators themselves explain exactly [why](https://golang.org/doc/faq#creating_a_new_language).
This is also a great video where Sameer Ajmani from Google explains the need for [Golang](https://youtu.be/5bYO60-qYOI?t=3m6s)

## Golang in short
Golang is a fairly new language built from the ground up, developed by Google and released in 2009.  
Go is somewhere between the compiled languages as C and C++, and the interpreted languages as Java. Go is a compiled language and outputs only one single binary file. But when working on a Go project you can also run the program as an interpreted language without having to compile first.  
Go also has object oriented-like features but uses them in a slighty different way. Go has no classes nor inheritence but uses alternative mechanisms to define relationships.  
### Code  
Go code is easy to write, clean and minimalistic and ```gofmt``` formats the code for you so every code looks the same. This makes working on large projects with lots of developers al lot easier.
### Packages  
Packages in Go and are very complete, powerful and versatile. They will save you a lot of time when, for instance, implementing network protocols, hashes and IO operations.
### Concurrency 
One of the great features in Golang is concurrency. In Go it's fairly easy to write a program running multiple processes at the same time. The combination of concurrency and the powerful packages is the reason why programs as Docker and Kubernetes are written in Go.

## Install Golang  
Go to the official [website](https://golang.org) and follow the install instructions.  
Ubuntu users can follow this [guid](https://tecadmin.net/install-go-on-ubuntu/)
Make you set the GOPATH correctly according to your OS.  
To unlock the full potential of Spacemacs and Golang, follow these steps:
1. Add these lines to your .bash_profile or .bashrc depending on OS:
```bash  
export GOPATH=$HOME/go # or the path you want Go to live in.
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```  

2. Make sure your GOPATH has these three folders:
  * bin
  * pkg
  * src

3. When you use Spacemacs as your editor, you can use life-changing enhcancements as godef, go-eldoc and guru, install these packages:
```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v golang.org/x/tools/cmd/goimports
```
Enabling goimports, open the `.spacemacs` file and look for:
`(defun dotspacemacs/user-init ()`  
Then add:
`(setq gofmt-command "goimports")`   
To automatically use syntax highlighting, add this line:
`(setq-default go-guru-hl-identifier-mode t)`  
So the code looks like this:
  ```cl  
  (defun dotspacemacs/user-init ()
    (setq gofmt-command "goimports")
    (setq-default go-guru-hl-identifier-mode t)
    "Initialization function for user code.
  It is called immediately after `dotspacemacs/init', before layer configuration
  executes.
  This function is mostly useful for variables that need to be set
  before packages are loaded. If you are unsure, you should try in setting them in
  `dotspacemacs/user-config' first."
  )
  ```  
  
4. In order automatically use the highlighting in every buffer, add the `(add-hook 'go-mode-hook #'go-guru-hl-identifier-mode)` below `(defun dotspacemacs/user-config ()` so it looks like this:
```cl  
  (defun dotspacemacs/user-config ()
    (add-hook 'go-mode-hook #'go-guru-hl-identifier-mode)
    "Configuration function for user code.
  This function is called at the very end of Spacemacs initialization after
  layers configuration.
  This is the place where most of your configurations should be done. Unless it is
  explicitly specified that a variable should be set before a package is loaded,
  you should place your code here."
    )
```  

5. In order to reload Spacemacs, press `SPC f e R`  
6. Now open Spacemacs and create a .go file byt typing `SPC f f` and typing in your filename plus .go extension.
Spacemacs will now import the go layer and is able to use all the installed packages.   
To get a grasp of how these functions work, you an read this [document](https://docs.google.com/document/d/1_Y9xCEMj5S-7rv2ooHpZNH15JgRT5iM742gJkw5LtmQ/edit) and watch this [video](https://www.youtube.com/watch?v=ak97oH0D6fI)

**Update october 2018**  
Some guidelines to install Spacemacs on Linux.  
Install Emacs:  
`apt install emacs` (sudo when needed or other package manager for your system.)  
Check if Emacs is 24.4 or above:  
`emacs --version`  
If Emacs is below 24.4:  
`apt install emacs25`  
Read the Spacemacs install instructions [here](https://github.com/syl20bnr/spacemacs).  
In short:  
* Clone the repository in the right location: `git clone https://github.com/syl20bnr/spacemacs ~/.emacs.d`
* Launch Emacs: `emacs` or when elpa needs an insecure http connection: `emacs --insecure`  
Spacemacs will be automaticly installed. Restart Emacs to complete the installation.  
When creating a .go file, Spacemacs will ask to install the Go-layer. Press yes.  

**Make sure you set your paths right otherwise these plugins below won't work!**    
Put this in your .profile / .bash_profile file:  
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin:/home/<user>/go/bin
```
> Somehow Linux doesn't always like environmental vars ($GOPATH and$GOROOT) as part of the PATH declaration so use absolute paths.  

To unlock some awesome tools in Go, read [this](https://github.com/syl20bnr/spacemacs/tree/master/layers/%2Blang/go) website.  
In short:
```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v golang.org/x/tools/cmd/goimports
```
If you want to use gometalinter:  
```
go get -u -v github.com/alecthomas/gometalinter
gometalinter --install --update
```
Only thing left is to activate some layers.  
Open the .spacemacs file in the .emacs.d folder or press `SPC f e d` in Emacs.  
Look for `dotspacemacs-configuration-layers` and un-comment `auto-completion` and `syntax-chacking`.  
You're probably set up now. Reboot your machine to make sure everything is loaded property.  

If the autocompletion isn't working, you need to enable auto-complete-mode. 
Open the .spacemacs file `SPC SPC`, enable auto-complete-mode and restart spacemacs.

## Learning go
### These are some useful introductions to Go:
* A Tour of [Go](https://tour.golang.org/) by Google.  
* Safaribook Online - Introduction to [Go](https://www.safaribooksonline.com/library/view/introducing-go/9781491941997/preface01.html#idp72384)  
  For this course you need to login. Ask us for the proper login information.  
* Miek Gieben has made a nice book  about [Learning Go](https://miek.nl/downloads/2015/go.pdf)(pdf warning).  
* Caleb Doxsey wrote this introduction for Go both on the [web](https://www.golang-book.com/books/intro) as a [PDF](https://www.golang-book.com/public/pdf/gobook.0.pdf)
* Alan Donovan and Brian Kernighan wrote [The Go Programming Language](http://www.gopl.io/). 

### Additional resources
* Information for all the packages in Go from the official [website](https://golang.org/pkg/)
* Dave Cheney has a nice [page](https://dave.cheney.net/resources-for-new-go-programmers) with valuable resources for Go.
* A curated [list](https://awesome-go.com/) of awesome Go frameworks, libraries and software.
* A series of [walkthroughs ](https://medium.com/go-walkthrough)to help you understand the Go standard library better.

### Nice excercises  
Test your skills and see if you can crack these [excersises](https://exercism.io/tracks/go/exercises).  
See also the various answers by other users.   

### Interesting Go talks
* [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
  A basic (and quick) introduction (with examples) for concurrent programming in Go. Very usefull for getting a quick overview of concurrent go programming.
* [Google I/O 2013 - Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)
  As the title describes, this talks goes more in depth about concurrent programming in Go. This too is a video that show a very quich but detailed overview.
* [Effectivity of Go](https://www.youtube.com/watch?v=cQ7STILAS0M) explained by one of the creators

### Use the power
Go has a few features that and packages that make Go so powerful and easy to use. It's worth paying extra attention to the items in the list below.   
* Interfaces  
You can use interfaces in a few different ways to really boost your code. Read this [post](http://spf13.com/post/is-go-object-oriented/) by Steve Francia about structs, methods and interfaces and how it relates to OOP.
* Goroutines and channels  
Concurrency in Go is powerful when knowing how to use it.
This is a great [video](https://youtu.be/f6kdp27TYZs) where Rob Pike explains Go channels, goroutines and patterns with example code.  
Divan has made concurrency [visual](http://divan.github.io/posts/go_concurrency_visualize/) for you and explains this in a [video](https://www.youtube.com/watch?v=KyuFeiG3Y60).  
Tutorial about concurrency and [wait-groups](https://medium.com/@matryer/very-basic-concurrency-for-beginners-in-go-663e63c6ba07)
* Packages
You don't have to re-invent the wheel. Go has core packages that so you don't have to build everything from scratch. Learning how to use them will make your life a lot easier:  
  * Strings package: [strings](https://golang.org/pkg/strings/) 
  * Input & output: [io package](https://golang.org/pkg/io/) 
  * Files & folders [os package](https://golang.org/pkg/os/) 
  * Errors [error package](https://golang.org/pkg/errors/)
  * Hashes & cryptography [crypto package](https://golang.org/pkg/crypto/)
  * Servers: [net package](https://golang.org/pkg/net/). This is where the power of Golang comes in. Networking is made very easy and the net package has built in protocols as http, tcp and rcp.   

### Testing
Go has a testing [package](https://golang.org/pkg/testing/) package.   
Video on [testing](https://www.youtube.com/watch?v=ndmB0bj7eyw) in Go.  

### Reference
The complete [Go Programming Language Specification](https://golang.org/ref/spec)  
Tutorial on [for loops](http://golangtutorials.blogspot.com/2011/06/control-structures-go-for-loop-break.html)


## Must read stuff
* A string is not a [string](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
* Blogpost about [strings](https://blog.golang.org/strings) in Go.
* Strings to bytes and [why](https://medium.com/go-walkthrough/go-walkthrough-bytes-strings-packages-499be9f4b5bd)
* Great explanation on how Go appends to slices for you and what happens [internally](https://blog.golang.org/go-slices-usage-and-internals)
* Getting a grasp on [pointers](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)

## Protocols  
Learn about all the different protocols you'll probably going to work with.
### HTTP
Great explanation of the HTTP protocol and how it communicates can be found [here](https://www.tutorialspoint.com/http/index.htm) and [here](https://code.tutsplus.com/tutorials/http-the-protocol-every-web-developer-must-know-part-1--net-31177)  

### Websocket
WebSockets provide a persistent connection between a client and server that both parties can use to start sending data at any time.  
A websocket begins as a HTTP request but then upgrades to a websocket connection.  


## Basics  
* Understand [pipes](https://ryanstutorials.net/linuxtutorial/piping.php) STDIN / STDOUT / STDERR about 

## Go and Docker
* Nice [blogpost](https://blog.docker.com/2016/09/docker-golang/) how to use Go and Docker.  
* Build [standalone](https://hub.docker.com/_/golang/) Docker images that run Go.   
* Complete [repository](https://github.com/docker-library/golang) with different Go versions and base images.  

## Dep   

Dep is a dependency management tool for the Go language. With `go get` you can install packages (imports) in your GOPATH folder so that they are available for every project.
However, when you are working on lots of different project, you don't always want to install all the packages. That's where Dep comes in.  
Dep allows you to install and maintain the packages (and which versions) you need just for that project.  
Check the dep [website](https://golang.github.io/dep/) for how to use dep.  

