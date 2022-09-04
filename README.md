# Build And Deploy to AWS EKS using GitHub Actions

Simple example to build and push Go and NodeJS image to DockerHub and deploy it to Amazon Elastic Kubernetes Service (EKS).

## Description

Test all codes before build and push Go and NodeJS images to DockerHub. Build process will not be executed if test process is failing. Deployment to AWS EKS will be executed right after build process finished. Approval needed to continue the deployment.

## Getting Started

### Dependencies

* Go v1.14.2
* Node v16.13.1

### Prerequisites

* Put all your Go scripts in go folder
* Put all your NodeJS scripts in node-js/src folder
* Set the repository secrets in Settings > Secrets > Actions
* List of secrets
  | Name       |
  | ------------- |
  | AWS_ACCESS_KEY_ID      |
  | AWS_SECRET_ACCESS_KEY      |
  | EKS_CLUSTER |
  | EKS_REGION|
  | DOCKERHUB_USERNAME|
  | DOCKERHUB_TOKEN |

### Usage

* Deployment will be executed automatically if there is any changes pushed to **main** branch 
* Deployment can be executed manually from **Actions** > **Test - Build - Deploy** > **Run Workflow** 
  
  
## Acknowledgments

Inspiration, code snippets, etc.
* [testing-go](https://blog.questionable.services/article/testing-http-handlers-go/)
* [simple-go-app](https://github.com/gscho/simple-go-app)
* [simple-node-js-app](https://github.com/ashleydavis/docker-nodejs-examples)
* [helm-charts](https://helm.sh/)
