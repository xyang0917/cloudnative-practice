import java.text.SimpleDateFormat;
	
pipeline {
  agent {
    kubernetes {
      yaml """
            apiVersion: v1
            kind: Pod
            metadata:
              labels:
                jenkins: worker
              namespace: jenkins
            spec:
              containers:
              - name: kaniko
                image: gcr.io/kaniko-project/executor:debug
                command:
                - sleep
                args:
                - 99999
                tty: true
                volumeMounts:
                  - name: docker-secret
                    mountPath: /kaniko/.docker
                    readOnly: true
              volumes:
              - name: docker-secret
                secret:
                  secretName: docker-aliyun-acr
                  items:
                    - key: .dockerconfigjson
                      path: config.json

            """
    }
  }

  environment {
    DATED_GIT_HASH = "${new SimpleDateFormat("yyMMddHHmmss").format(new Date())}${GIT_COMMIT.take(6)}"
  }

  stages {
    stage('Configure') {
      steps {
        echo "hello, starting"
      }
    }

    stage('Build with Kaniko') {
      steps {
        container('kaniko') {
          sh '/kaniko/executor -f `pwd`/httpserver/Dockerfile -c `pwd`/httpserver --cache=true \
                  --destination=registry.cn-hongkong.aliyuncs.com/yangxin/httpserver:${DATED_GIT_HASH} \
                  --insecure \
                  --skip-tls-verify  \
                  -v=debug'
        }
      }
    }  
  }
}
