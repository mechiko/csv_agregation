# set shell := ["pwsh", "", "-CommandWithArgs"]
# set positional-arguments
shebang := 'pwsh.exe'
# Variables
exe_name := "agregat"
mod_name := "agregat"
ld_flags :="-s -w"

default:
  just --list

win64:
    #!{{shebang}}
    $env:Path = "C:\Go\go.125\bin;C:\go\gcc\mingw64\bin;" + $env:Path
    $env:GOARCH = "amd64"
    $env:GOOS = "windows"
    $env:CGO_ENABLED = 1
    if (-Not (Test-Path go.mod)) {
      go mod init {{mod_name}}
    }
    go mod tidy -v
    if(-Not $?) { exit }
    Remove-Item .\.dist\{{exe_name}}.exe, .\.dist\{{exe_name}}_64.exe 2>$null
    go build -ldflags="{{ld_flags}}" -o ./.dist/{{exe_name}}_64.exe .
    if(-Not $?) { exit }
    upx --force-overwrite -o ./.dist/{{exe_name}}.exe ./.dist/{{exe_name}}_64.exe
