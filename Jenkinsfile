pipeline {
    agent any
    tools {
        go 'go 1.18'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                 echo 'Compiling and building'
                 sh 'go build'
            }
        }

        stage('Test') {
            steps {
                     echo 'Running testing 1'
                     sh 'cd delivery'
                     sh 'cd user_delivery && go test'
                }
            steps {
                     echo 'Running testing 2'
                     sh 'cd usecase'
                     sh 'cd user_usecase && go test'
                }
            steps {
                     echo 'Running testing 3'
                     sh 'cd repository'
                     sh 'cd user_repository && go test'
                }
            
        }
        
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}