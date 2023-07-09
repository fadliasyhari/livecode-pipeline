pipeline {
    agent any
    environment {
        GIT_URL = 'https://github.com/fadliasyhari/livecode-pipeline.git'
        BRANCH = 'main'
        IMAGE = 'my-golang-test'
        CONTAINER = 'my-golang-test-app'
        DOCKER_APP = 'docker'
        DB_HOST = 'product-db'
        DB_USER = 'postgres'
        DB_NAME = 'book_manaegement_system'
        DB_PASS = 'password'
        DB_PORT = '5432'
        API_PORT = '8000'
        JOB_NAME = 'livecode-pipeline'

    }
    stages {
        stage("Cleaning up") {
            steps {
                echo 'Cleaning up'
                sh "${DOCKER_APP} rm -f ${CONTAINER} || true"
            }
        }

        stage("Clone") {
            steps {
                echo 'Clone'
                git branch: "${BRANCH}", url: "${GIT_URL}"
            }
        }

        stage("Build and Run") {
            steps {
                echo 'Build and Run'
                sh "DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} DB_NAME=${DB_NAME} DB_USER=${DB_USER} DB_PASS=${DB_PASS} API_PORT=${API_PORT} ${DOCKER_APP} compose up -d"
            }
        }
    }
    post {
        always {
            emailext to: "asyharifadli@gmail.com",
            subject: "Livecode Pipeline Notification",
            body: "Pipeline ${env.JOB_NAME} running with ${env.GIT_COMMIT}",
            attachLog: true
            slackSend botUser: true, 
            channel: 'general',
            message: "Pipeline ${env.JOB_NAME} running with ${env.GIT_COMMIT}", 
            tokenCredentialId: 'Slack OAUTH'
        }
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
    }
}
