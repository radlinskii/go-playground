/*
Gowhich is an implementation of which(1) command in go.

Usage:
	gowhich [flags] path ...

The flags are:
	-s:		List all instances of executables found (instead of just the first one of each).
	-a: 	No output, just return 0 if all of the executables are found, or 1 if some were not found.

*/

package main
