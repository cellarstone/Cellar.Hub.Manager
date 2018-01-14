pipeline {
  agent any
  
  stages {
    stage('Pre Test') {
      steps {
        echo 'Testing Golang'
        sh 'export GOPATH=/home/kellersteinlukas/go'
        sh 'export PATH=$GOPATH/bin:$PATH'
        sh 'export GOROOT=/usr/lib/go-1.9'
        sh 'export PATH=$GOROOT/bin:$PATH'
        echo '$GOROOT'
        echo '$GOPATH'
        sh 'go version'
      }
    }
    stage('Build') {
      when {
        branch 'master' 
      }
      steps {
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
                --version="0.3.4" \
                --platforms="darwin_amd64 linux_amd64" \
                --signing-key=equinox.key \
                --app="app_h9SyPnPqLpq" \
                --token="fHeN81JECeiVAxoiJfEyPxBGSdMnBxVjsxZffG7wrHgEvwqJshuF" \
                ../'
        }
      }
    }
  }
}