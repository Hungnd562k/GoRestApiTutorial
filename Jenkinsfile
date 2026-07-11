pipeline {
    agent { label 'Agent-1' }
    environment {
        COMMIT_HASH = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
    }
    stages {
        stage('Clone source') {
            steps {
                echo 'Checking and updating source code...'
                // Kiểm tra nếu thư mục .git đã tồn tại
                sh '''
                    if [ -d ".git" ]; then
                        echo "Repository exists. Pulling latest changes..."
                        git fetch --all
                        git checkout master
                        git pull origin master
                    else
                        echo "Repository does not exist. Cloning from scratch..."
                        git clone https://github.com/Hungnd562k/GoRestApiTutorial.git .
                        git checkout master
                    fi
                '''
            }
        }
        stage('Build Image') {
            steps {
                sh "docker build -t hungnd2/go-rest-api-turtorial:${COMMIT_HASH} ."
            }
        }
        stage('Push Image') {
            steps {
                // Đã loại bỏ dấu chấm (.) dư thừa ở cuối lệnh push
                sh "docker push hungnd2/go-rest-api-turtorial:${COMMIT_HASH}"
            }
        }
    }
}