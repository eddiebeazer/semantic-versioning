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
                bat 'powershell.exe -file .\\build.ps1'
            }
        }
    }
}
