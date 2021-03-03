vectorworks-app-cleaner

This software will remove files on your harddrive without confirmation.  Please understand what is being done before running this software.

## The Software

This software is made specifically for the Vectorworks application suite and used in order to complete otherwise time consuming tasks

## Windows:
Windows requires the following dlls: `webview.dll` `WebView2Loader.dll`
Build: `go build -ldflags="-H windowsgui" -o webview-example.exe`

```cmd
# Windows needs this line run in order to use the application
CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"
```

There is also a package script that will place everything within the build/ci folder.

This workflow will be optimized in the coming days.

This repo requires go v1.16+