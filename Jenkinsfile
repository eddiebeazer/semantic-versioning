pipeline {
    agent any
    tools {
        go '1.18'
    }
    stages {
        stage('Installing Dependencies') {
            steps {
                echo 'Getting modules'
                bat 'go get -u -d ./...'
            }
        }
        stage('Testing') {
            steps {
                golangTesting()
            }
        }
        stage('Build') {
            steps {
                powershell "$Env:GOOS = 'windows'; $Env:GOARCH = 'amd64'; go build -o .\\bin\\semantic-versioning.exe"
                powershell "$Env:GOOS = 'linux'; $Env:GOARCH = 'amd64'; go build -o .\\bin\\semantic-versioning"
            }
        }
    }
}
