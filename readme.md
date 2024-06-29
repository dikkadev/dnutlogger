# dnutlogger

### **D**o **N**ot **U**se **T**his **Logger** (Seriously, Why Are You Even Here?)

---

Oh, you found dnutlogger. Guess you're out of good options. This is what I use for logging because I can't be bothered with anything better. If you're here, you must share my questionable standards. Let's get on with it.

## What is dnutlogger?

Glad you asked. It's a logger. That's it. It logs stuff. To the console. Wow. Revolutionary, I know. But hey, it works for me, so it might just work for you. Or not. Probably not. But who cares?

## How to Install

You really want to do this, huh? Fine. Here you go:

```sh
go get github.com/sett17/dnutlogger
```

## Usage

Alright, so you’ve gone ahead and installed it. Bold move. Here’s how you can use this masterpiece of logging technology:

```go
package main

import (
    "errors"
    log "github.com/sett17/dnutlogger"
)

func main() {
    // Optional configurations
    log.SetMinLevel(log.DEBUG) // Set the minimum log level to DEBUG
    log.UseColor(false)        // Disable colored output (default is true)

    // Logging a debug message
    log.Debug("This is a debug message")

    // Logging a debug formatted message
    log.Debugf("This is a debug message with a number: %d", 42)

    // Logging an info message
    log.Info("This is an info message")

    // Logging an info formatted message
    log.Infof("This is an info message with a string: %s", "example")

    // Logging a warning message
    log.Warn("This is a warning message")

    // Logging a warning formatted message
    log.Warnf("This is a warning message with a float: %.2f", 3.14)

    // Logging a success message
    log.Success("This is a success message")

    // Logging a success formatted message
    log.Successf("This is a success message with a boolean: %t", true)

    // Logging an error message without exiting
    log.Error(false, "This is an error message without exiting")

    // Logging an error formatted message without exiting
    log.Errorf(false, "This is an error message with a struct: %+v", struct{ Name string }{"example"})

    // Logging an error with a stack trace and exiting
    log.Err(true, errors.New("This is a fatal error"))
}
```

## Log Levels

Wow, you actually care about the log levels? Fine, here they are:

- **DEBUG**: For when you need to log every single step because you have no idea what's going on.
- **INFO**: Standard log level, because you’re supposed to care about these messages.
- **SUCCESS**: Everything went better than expected.
- **WARN**: Things are getting sketchy, might want to pay attention.
- **ERROR**: You messed up. Big time.

_This is also the order used with `SetMinLevel()`, so setting the minimum log level to `SUCCESS` will log everything from `SUCCESS` to `ERROR`._

## Configuration

### Set Minimum Log Level

Because obviously, you need to filter out the noise:

```go
log.SetMinLevel(log.WARN)
```

### Enable/Disable Color

Make your console pretty, or not:

```go
log.UseColor(true) // or false if you're boring
```

## License

Do whatever you want. I do not care. Seriously.

---

Well, there you have it. You’ve just spent time reading a README for a logger that I don’t even recommend you use. But you do you. Enjoy (or not).
