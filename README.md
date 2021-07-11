# windows-shortcuts-golang

This project is a test for using the Win32 api in Golang to get the currently pressed keys, and detect if there is a combination of keystrokes associated to a command (ctrl, shift or alt). It's not intended to be used in production. 

To clone this project. 

```console
git clone https://github.com/DatsGabs/windows-shortcuts-golang
cd windows-shortcuts-golang
go run main.go 
```

You can use go build tot build an exectutable. 

```console
go build
./shortcuts.exe
```

If you want to run the application without the prompt you can follow [this tutorial my blog](https://gabriellazcano.com/blog/how-to-run-a-exe-without-opening-a-prompt-in-windows/)
