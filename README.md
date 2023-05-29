# CI/CD for Amazon EKS using Github Action

## Descriptions
This repo contains two backend applications named **rentcar** using on Go programming language and **carshop** using on Node.js that will automatically deployed to Amazon EKS 
cluster using CI/CD Github Action.

## Workflows
The flows of CI/CD GitHub in `.github/workflows` folder are:
- First, the runner will checkout the code in each of folder repo (e.g if there is any update on carshop folder by push, it will trigger the ci/cd)
- Then, configure the AWS credentials to login to the AWS env using secrets stored in the Action secret
- Login to Amazon ECR (Elastic Container Registry) using the AWS credentials before
- Build, tag, and push each image application to Amazon ECR
- Deploy the backend to Amazon EKS using a newly pushed image
- Verify the deployment

## rentcar
This is a simple rest api for rent car service based on Go programming language without relying on external frameworks. The folder contains code itself and manifests file on 
manifests folder for Kubernetes deployment and service. Source code: [https://github.com/worlpaker/go_rentacar_restapi](https://github.com/worlpaker/go_rentacar_restapi)

### Features
- Manual handling of routing, middlewares, and helper functions.
- Swagger Documentation
- Go unit tests with mocking MongoDB

### Components
- Go
- Javascript
- Docker
- MongoDB

### API Endpoints
- https://rentcar.dwiananda.click/ (404 not found)
- Swagger: [https://rentcar.dwiananda.click/api/swagger/index.html](https://rentcar.dwiananda.click/api/swagger/index.html)
- cars
  - **GET** : [https://rentcar.dwiananda.click/api/cars/available](https://rentcar.dwiananda.click/api/cars/available)
  - **POST** : [https://rentcar.dwiananda.click/api/cars/reserve](https://rentcar.dwiananda.click/api/cars/reserve)
    <br> Example value: 
    ```json
    {
      "name": "string",
      "nation_id": "string",
      "phone_number": "string",
      "surname": "string"
    }
    ```
  - **GET** : [https://rentcar.dwiananda.click/api/cars/showreservedcars](https://rentcar.dwiananda.click/api/cars/showreservedcars)
- Locations
  - **GET** : [https://rentcar.dwiananda.click/api/locations/show](https://rentcar.dwiananda.click/api/locations/show)


## carshop
This is a simple rest api for car shop service based on Node.js and MongoDB as a database. The folder contains code itself and manifests file on manifests folder for Kubernetes 
deployment and service. Source code: [https://github.com/jonathankarlinski/car-shop](https://github.com/jonathankarlinski/car-shop)

### Components
- Typescript
- MongoDB
- Mongoose
- NodeJS
- Docker
- Mocha
- Chai
- Sinon
- POO
- CRUD

### API Endpoints
- https://carshop.dwiananda.click/ (404 not found)
- cars
  - **GET** : [https://carshop.dwiananda.click/cars](https://carshop.dwiananda.click/cars)
  - **POST** : [https://carshop.dwiananda.click/cars](https://carshop.dwiananda.click/cars)
    <br> Example value: 
    ```json
    {
      "model": "Marea",
      "year": 2002,
      "color": "Black",
      "status": true,
      "buyValue": 15.990,
      "doorsQty": 4,
      "seatsQty": 5
    }
    ```
- motorcycles
  - **GET** : [https://carshop.dwiananda.click/motorcycles](https://carshop.dwiananda.click/motorcycles)
  - **POST** : [https://carshop.dwiananda.click/motorcycles](https://carshop.dwiananda.click/motorcycles)
  <br> Example value:
  ```json
  {
    "id": "6348513f34c397abcad040b2",
    "model": "Honda Cb 600f Hornet",
    "year": 2005,
    "color": "Yellow",
    "status": true,
    "buyValue": 30.000,
    "category": "Street",
    "engineCapacity": 600
  }
  ```

## Infrastructure Stacks
- **Amazon EKS**
  <br> This service is for the Kubernetes cluster platform managed by AWS. I've deployed using `eksctl` tool with cloudformation stack.
- **Amazon EC2 (Elastic Compute Cloud)**
  <br> This service is for the worker node of the Kubernetes cluster. 
- **Amazon VPC (Virtual Private Cloud)**
  <br> Networking service for most of all resources.
- **Amazon Route 53**
  <br> This service is for managing the DNS server including its record. 
- **Amazon ECR (Elastic Container Registry)**
  <br> This service is to store rentcar and carshop images.
- **AWS Certificate Manager (ACM)**
  <br> This service is for provision manage certificates of applications
- **ELB using Application Load Balancer (ALB)**
  <br> This service is a load balancer for the application deployed on Amazon EKS
- **AWS CloudFormation**
  <br> I've deployed Amazon EKS using a template that runs on cloudformation
- **AWS IAM (Identity and Access Management)**
  <br> This service is for configuring and managing access between the resources (e.g: access push the image from GitHub Actions to Amazon ECR)
- **GitHub & GitHub Actions**
- NB: **Screenshots for documentation attached on [images](images/) folder**

## References
- https://docs.aws.amazon.com/
- https://github.com/marketplace/actions/kubectl-aws-eks
- https://github.com/aws-actions/amazon-ecr-login
- https://github.com/aws-actions/configure-aws-credentials
- https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions
- https://docs.github.com/en/actions/security-guides/encrypted-secrets
- https://earthly.dev/blog/aws-ecr-github-actions/
- https://docs.github.com/en/actions/deployment/deploying-to-your-cloud-provider/deploying-to-amazon-elastic-container-service
