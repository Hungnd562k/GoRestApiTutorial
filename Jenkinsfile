pipeline {
    agent { label 'Agent-1' }
    environment {
        // Định nghĩa biến global tại đây
        COMMIT_HASH = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
    }
    stages {
        stage('Clone source') {
            steps {
                echo 'Cloning source code...'
                sh 'git clone https://github.com/Hungnd562k/GoRestApiTutorial.git'
                sh 'git checkout master'
            }
        }
        stage('Build Image') {
            steps {
                sh "docker build -t hungnd2/go-rest-api-turtorial:${COMMIT_HASH} ."
            }
        }
        stage('Push Image') {
            steps {
                sh "docker push -t hungnd2/go-rest-api-turtorial:${COMMIT_HASH} ."
            }
        }
    }
}
