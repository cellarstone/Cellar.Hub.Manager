pipeline {
  agent any
  
  stages {
    stage('Equinox') {
      when {
        branch 'master' 
      }
      steps {
        parallel (
          manager: {
            sh './equinox release \
                    --version="0.3.4" \
                    --platforms="darwin_amd64 linux_amd64" \
                    --signing-key=equinox.key \
                    --app="app_h9SyPnPqLpq" \
                    --token="fHeN81JECeiVAxoiJfEyPxBGSdMnBxVjsxZffG7wrHgEvwqJshuF" \
                    ../'
          }
        )
      }
    }
  }
}