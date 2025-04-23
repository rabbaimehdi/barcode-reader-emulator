rsrc_windows_386.exe -manifest "main.manifest" -o "rsrc.syso" -ico "app.ico"
go build -ldflags -H=windowsgui