/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.26.4-alpine3.24' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}