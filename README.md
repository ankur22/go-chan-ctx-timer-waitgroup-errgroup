# Code for article on medium.com

## ctxwithcancel.go

`go run ctxwithcancel.go`

This shows how to use the context done channel to wait for when the done channel is explicitly closed by SIGINT (ctrl+c).

## chanok.go

`go run chanok.go`

Using the ok variable read from the channel to identify when the channel was closed so we can shutdown the sender goroutine. The Channel is closed when the reader has completed reading and the channel is closed (deferred).

## waitgrp.go

`go run waitgrp.go`

Uses WaitGroup to wait for the goroutines to complete. Still uses the ok variable from the channel to determine when the channel is closed.

## errgrp.go

`go run errgrp.go`

Uses errgroup instead of WaitGroup. There are three ways that wait() will unblock:

1. All goroutines return nil error;
2. When at least one goroutines returns an error;
3. The parent context of the errgroup context is closed.

## timer.go

`go run timer.go`

This shows how to use a timer when waiting to read off a channel where the writer to the channel is slow. We may want to abort reading off a channel if it takes too long.

## range.go

`go run range.go`

This ranges over a channel. This means we don't need to explicitly check the ok variable returned from the channel for when the channel is closed. We can't use timers though, and we also need to remember to use the default case otherwise we would block on ctx.done().

## efficient.go

`go run efficient.go`

We should retrieve the done channel from the context before using it in the goroutine as this avoid the mutex lock and unlocks.  
Timers can be resued instead of creating a new one when it fires, just remember to rest the timer.
