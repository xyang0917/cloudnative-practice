pipeline {
  agent any
  stages {
    stage('Config') {
      steps {
        sh 'echo begin config'
      }
    }

    stage('Build with Kaniko') {
      steps {
        container(shell: 'sh \'/kaniko/executor -f `pwd`/Dockerfile -c `pwd`/src --cache=true \\ 	          --destination=cloudnative.azurecr.io/httpserver:${DATED_GIT_HASH} \\ 	                  --insecure \\ 	                  --skip-tls-verify  \\ 	                  -v=debug\'', name: 'kaniko')
      }
    }

  }
}