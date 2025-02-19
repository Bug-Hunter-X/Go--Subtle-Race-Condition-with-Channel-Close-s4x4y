# Go: Subtle Race Condition with Channel Close

This repository demonstrates a subtle race condition in a Go program involving channel closure.  The program appears to work correctly in simple tests, but it hides a potential issue that can cause problems in more complex concurrent scenarios. 

The `bug.go` file contains the buggy code, illustrating the race condition. The solution is provided in `bugSolution.go`, highlighting best practices for handling channel closure and synchronization in Go.