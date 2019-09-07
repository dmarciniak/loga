# loga (log-analizer)

[![Go Report Card](https://goreportcard.com/badge/github.com/dmarciniak/loga)](https://goreportcard.com/report/github.com/dmarciniak/loga)

Console UI application to read logs from files, sorting them and allows analize them.

## Uses
Application takes list of files:
```
./loga file1.log file2.log [...]
```
## Main features
* console user-friendly interface
* sorting loaded logs
* filtering them (regex)
* popup window to show all logs from one file (selected by current log line)
* coloured logs per file

## Screenshots

![preview1](https://user-images.githubusercontent.com/39051624/61814414-55b03380-ae48-11e9-9553-68855790a961.png)

![preview2](https://user-images.githubusercontent.com/39051624/61815992-b9882b80-ae4b-11e9-9361-51a516b9c63a.png)


## Note
This aplication uses GoCui library to display interface: https://github.com/jroimartin/gocui

