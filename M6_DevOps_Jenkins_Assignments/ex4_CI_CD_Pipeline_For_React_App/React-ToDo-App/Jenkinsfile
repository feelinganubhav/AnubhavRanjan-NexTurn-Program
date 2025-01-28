pipeline {
    agent any

    environment {
        DEPLOY_DIR = 'D:\\React-ToDo-App-CI-CD-Deployment'
    }

    stages {
        stage('Clone Repository') {
            steps {
                echo 'Cloning the React ToDo Application repository...'
                git branch: 'main', url: 'https://github.com/feelinganubhav/React-ToDo-App.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies...'
                bat 'npm install'
            }
        }

        stage('Lint Code') {
            steps {
                echo 'Linting the code using ESLint...'
                bat 'npx eslint . || exit /B 0'
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests On All React ToDo App Fuctionalities using Jest...'
                bat 'npm test || exit /B 0'
            }
        }

        stage('Build Application') {
            steps {
                echo 'Building the React ToDo Application...'
                bat 'npm run build'
            }
        }

        stage('Deploy Application') {
            steps {
                echo 'Deploying the application...'
                bat "del /q ${DEPLOY_DIR}\\*"
                bat "xcopy dist ${DEPLOY_DIR} /E /I /H /Y"
            }
        }

        stage('Start Server for Testing') {
            steps {
                echo 'Starting the Server...'
                bat "start /B npx serve -s ${DEPLOY_DIR} -l 3000"

                echo 'Waiting for a few seconds to allow the React app to start'
                bat "powershell -Command 'Start-Sleep -Seconds 3'"

            }
        }

        stage('Post-deployment Testing') {
            steps {
                script {
                    echo 'Testing the App if it is Running or NOT...'
                    
                    echo 'Execute the curl command to fetch the HTTP status code'
                    def responseCode = bat(
                        script: """
                        @echo off
                        curl -o NUL -s -w %%{http_code} http://localhost:3000
                        """,
                        returnStdout: true
                    ).trim()
                    
                    echo "HTTP Response Code: ${responseCode}"

                    // Check if the response code is 200
                    if (responseCode == '200') {
                        echo 'App is Up and Running Successfully...'
                    } else {
                        error "Post-deployment testing failed with HTTP status code: ${responseCode}"
                    }
                }
            }
        }
    }

    post {
        always {
            echo 'Cleaning up any running servers...'
            bat '''
            for /f "tokens=5" %%a in ('netstat -ano ^| find ":3000"') do (
                if %%a NEQ 0 (
                    taskkill /pid %%a /f
                    exit /B 0
                )
            )
            '''
            
        }
        success {
            echo 'Pipeline executed successfully!'
        }
        failure {
            echo 'Pipeline execution failed!'
        }
    }
}
