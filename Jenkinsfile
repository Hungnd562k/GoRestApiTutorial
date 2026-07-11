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
            environment {
                // Nạp Username và Password từ Credentials vào 2 biến môi trường
                DOCKER_HUB = credentials('docker-hub-credentials') 
            }
            steps {
                echo "Logging into Docker Hub..."
                // Lệnh đăng nhập sử dụng biến môi trường tự động sinh ra ($DOCKER_HUB_USR và $DOCKER_HUB_PSW)
                sh "echo \$DOCKER_HUB_PSW | docker login -u \$DOCKER_HUB_USR --password-stdin"
                
                echo "Pushing image..."
                sh "docker push hungnd2/go-rest-api-turtorial:${COMMIT_HASH}"
            }
        }
    }
}
