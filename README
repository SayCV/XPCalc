XPCalc
======

Based on the Walk library for the Go Programming Language.
Based on the windows XP calc.

Setup
=====

Make sure you have a working Go installation.
See [Getting Started](http://golang.org/doc/install.html)

##### Build app

In the root directory run

	go build
	
To get rid of the cmd window, instead run

	go build -ldflags="-H windowsgui"

##### Run app
	
	XPClac.exe

Application Manifest Files
==========================
Walk requires Common Controls 6. This means that you must put an appropriate
application manifest file either next to your executable or embedded as a
resource.

You can copy one of the application manifest files that come with the examples.

To embed a manifest file as a resource, you can use the [rsrc tool](https://github.com/akavel/rsrc).

IMPORTANT: If you don't embed a manifest as a resource, then you should not launch
your executable before the manifest file is in place.
If you do anyway, the program will not run properly. And worse, Windows will not
recognize a manifest file, you later drop next to the executable. To fix this,