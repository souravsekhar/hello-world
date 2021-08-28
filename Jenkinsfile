pipeline{
    agent none
    environment{
        // ECR_CREDS=credentials("ECR_CREDS")
        registry_url='aws_account_id.dkr.ecr.region.amazonaws.com/my-repository:tag'
    }
    stages{
        stage("Docker build and push image"){
            agent {label 'docker'}
            steps{
                script{
                    // sh 'echo $ECR_CREDS > keyfile.json'
                    sh 'docker login --username AWS --password-stdin aws_account_id.dkr.ecr.region.amazonaws.com'
                    sh 'docker build . -t ${registry_url}:latest'
                    sh 'docker push ${registry_url}:latest'
                }
            }
        }
    }
    post{
        failure{
            script{
                mail (to: sekhar.sourav8@gmail.com,
                subject: 'Job failed',
                body: 'Hello world app build and push failure')
            }
        }
    }
}