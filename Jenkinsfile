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
                 sh 'go mod vendor'
                 sh 'go build -mod=readonly'
            }
        }

        stage('Test') {
            steps {
                     echo 'Running testing 1'
                     sh 'cd delivery/user_delivery && go test'
                     sh 'cd ../..'
                     echo 'Running testing 2'
                     sh 'cd usecase/user_usecase && go test'
                     sh 'cd ../..'
                     echo 'Running testing 3'
                    //  sh 'cd repository/user_repository && go test'
                    //  sh 'cd ../..'
                    // cant test local repo on jenkins
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