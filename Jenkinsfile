pipeline {
  agent any
  triggers {
    githubPush()
  }
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
  tools {
    go 'Go 1.22.1'
  }
}
