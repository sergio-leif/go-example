pipeline {

  agent any

  stages {
    
      stage("Build image") {
            steps {
                script {
                    myapp = docker.build("sleivagc/go-web-app:2.0")
                }
            }
        }
    
      stage("Push image") {
            steps {
                script {
                    docker.withRegistry('', 'dockerhub-personal') {
                            myapp.push("latest")
                            myapp.push("2.0")
                    }
                }
            }
        }

  }

}
