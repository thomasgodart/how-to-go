
The goal of this **repository** is to teach you the "**way to Go**" with concrete examples and good practices. It could be a starting point for a company who wants to explore the possibilities of the Go language, or a support for a quick presentation of the language, or save there apps templates.

Using it needs:
* being introduced first to the [tour of Go](https://tour.golang.org/)
* having some sense of the [`pkg` repository](https://golang.org/pkg/)
* the target environment is [Ubuntu, Linux](https://ubuntu.com/)

It is composed of:
* 0.[`ini`](0-ini), the **documentation**
	* 0.**install** instructions
	* 1._**IDE setup**_
* 1.[`cmd`](1-cmd), the **command line** apps who quit when the work is done
	* 0.**empty**
		* the empty app is interesting to measure the Go language overhead
	* 1.**hello**
		* a classic, but will teach the command line options and writing styles
	* 2.**environment**
		* setup an environment to configure your app, overwrite it with the command line
* 2.[`srv`](2-srv), the **services** apps who stay alive and never quit
	* 0.**hello**
		* 0.**default**
			* the simplest "hello world" service
		* 1.**mux**
			* same "hello world" but using the famous `gorilla/mux` package
	* 1.**simple**
		* a simple web server with HTML templates and static directories
	* 2.**system**, a web server started and controlled by `systemd`
		* 0.**notify**
			* a simple server that answers to *readiness* and *liveness*
		* 1.**socket**
			* a more robust server that lets `systemd` deal with sockets
	* 3.**crud**, a simple "create retrieve update delete" app
		* 0.**file**, on a file system
		* 1.**sql**, on SQL using [Gorm](https://gorm.io/docs/)

Numbers are used in the directory names for ordering them.
