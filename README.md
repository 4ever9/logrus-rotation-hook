# Logrus Rotation Hook

## Install

```bash
go get github.com/4ever9/logrus-rotation-hook
```

## Usage

```go
// function `WithXXX` is optional.
hook, err := NewHook(
    WithMaxBackups(0),  // default is 0
    WithMaxSize(10),    // default is 20MB
    WithMaxAge(1),      // default is 1
    WithCompress(true), // default is false
    WithFilename("./logs/app.log") // default is ./app.log
)

if err != nil {
    return err
}

logrus.AddHook(hook)
```
