
The goal of this **repository** is to teach you the "**way to Go**" with concrete examples and good practices. It could be a starting point for a company who wants to explore the possibilities of the Go language, or a support for a quick presentation of the language, or save there apps templates. Using it needs:
* being introduced first to the [tour of Go](https://tour.golang.org/)
* having some sense of the [`pkg` repository](https://golang.org/pkg/)
* the target environment is [Ubuntu, Linux](https://ubuntu.com/)

It is composed of:
1. `ini`: the **documentation**
	1. the **install** instructions
	1. the **IDE setup**
1. `cmd`: the **command line** apps who quit when the work is done
	1. **empty**
		* the empty app is interesting to measure the Go language overhead
	1. **hello**
		* a classic, but will teach the command line options and writing styles
	1. **environment**
		* setup an environment to configure your app, overwrite it with the command line
1. `srv`: the **services** apps who stay alive and never quit
	1. **hello**
		* **default**
			* the most simple hello world service
		* **mux**
			* same hello world but using the famous `gorilla/mux` package
	1. **simple**
		* a simple web server with HTML templates
	1. **system**
		* a web server started and controlled by `systemd` (Ubuntu, Linux)
	1. **crud**, a simple "create retrieve update delete" app
		1. **file** system
		1. **SQL** using [Gorm](https://gorm.io/docs/)

Numbers are used in the directory names for ordering them.
