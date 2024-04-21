pipeline {
  agent any

  tools {go '1.22'}
  
  stages {
    stage('Build') {
      steps {
        sh 'go version'
        sh 'go build'
      }
    }

    stage('Test') {
      steps {
        sh 'go test'
      }
    }

  }
}
