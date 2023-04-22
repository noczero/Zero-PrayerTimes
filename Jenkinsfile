pipeline {
    agent any
    tools {
        go 'go-1.20'
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[credentialsId: 'github-credentials', url: 'git@github.com:noczero/Zero-PrayerTimes.git']])
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o main ./cmd/main.go'
            }
        }
        stage("Run") {
            steps {
                sh './main'
            }
        }
    }
}
