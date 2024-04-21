pipeline {
  agent any
  stages {
    stage('Checkout') {
      steps {
        git(url: 'https://github.com/olaung/go-app.git', branch: 'master')
      }
    }

    stage('Build') {
      steps {
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