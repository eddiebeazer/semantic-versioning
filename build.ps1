$Env:GOOS = "windows"; $Env:GOARCH = "amd64"

go build -o .\bin\semantic-versioning.exe

$Env:GOOS = "linux"; $Env:GOARCH = "amd64"

go build -o .\bin\semantic-versioning
