pipeline {
    agent any
    stages {
        stage('Build Image') {
            withEnv([
                "DOCKER_HOST=tcp://docker:2376",
                "DOCKER_CERT_PATH=/certs/client",
                "DOCKER_TLS_VERIFY=1"
            ]) {
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
            steps {
                sh 'docker rmi hungnd2/go-rest-api-turtorial:v2.0.0'
            }
        }
        stage('Ultimately final stage') {
            steps {
                echo 'This is the real final stage!'
            }
        }
    }
}