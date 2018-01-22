pipeline {
  agent any

  environment {
    GOPATH = '/var/lib/jenkins/go'
    GOROOT = '/usr/lib/go-1.9'
    PATH = "$PATH:$GOROOT/bin:$GOPATH"
  }
  
  stages {
    stage('Pre Test') {
      steps {
        echo 'Testing Golang'
        sh 'printenv'
        sh 'go version'
      }
    }
    stage('Build') {
      when {
        branch 'master' 
      }
      steps {
        sh 'go get github.com/gorilla/mux'
        sh 'go get github.com/equinox-io/equinox'
        sh 'go get github.com/facebookgo/grace/gracehttp'
        sh 'go get github.com/arschles/go-bindata-html-template'
        sh 'go get gopkg.in/resty.v1'
        sh 'go get cloud.google.com/go/storage'
        sh 'go get cloud.google.com/go/pubsub'
        sh 'go get golang.org/x/net/context'
        sh 'go get github.com/jaypipes/ghw'
        sh 'go-bindata views/...'
      }
    }
    stage('Equinox') {
      when {
        branch 'master' 
      }
      steps {
        dir ('equinox') { 
          sh './equinox release \
                --version="0.5.1" \
                --platforms="linux_amd64" \
                --signing-key=equinox.key \
                --app="app_h9SyPnPqLpq" \
                --token="fHeN81JECeiVAxoiJfEyPxBGSdMnBxVjsxZffG7wrHgEvwqJshuF" \
                ../'
        }
      }
    }
  }
}