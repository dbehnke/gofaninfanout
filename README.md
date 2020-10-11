# gofaninfanout
goofing around with go - fan in and fan out with go routines and channels


## goals

The goal of this project is to just simulate a work queue.

For the work queue, it will simply add two numbers and set a result.  

It will have a random delay during the work process to simulate variable work.

There will be limited number of runners, use channels to send work.

Use wait group to keep track of completed work.
