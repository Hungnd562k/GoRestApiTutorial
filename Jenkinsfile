pipeline {
    agent any
    stages {
        stage('Build Image') {
            steps {
                sh 'docker build -t hungnd2/go-rest-api-turtorial:v2.0.0 .'
            }
        }
        stage('Push Image') {
            steps {
                echo 'This is stage 2!'
            }
        }
        stage('Stage 3') {
            steps {
                echo 'This is the final stage!'
            }
        }
        stage('Remove Image after push to Registry') {
            sh 'docker rm hungnd2/go-rest-api-turtorial:v2.0.0'
        }
        stage('Ultimately final stage') {
            steps {
                echo 'This is the real final stage!'
            }
        }
    }
}