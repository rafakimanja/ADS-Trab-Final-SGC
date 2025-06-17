pipeline {
    agent any
    tools {
        go 'Go 1.24'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
       
        stage("build") {
            steps {
                sh 'go run main.go'
            }
        }

    }
}