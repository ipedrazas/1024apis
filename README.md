# 1024 APIs
[![wercker status](https://app.wercker.com/status/4072de4970ad81233c00449ba12ef4eb/s/master "wercker status")](https://app.wercker.com/project/byKey/4072de4970ad81233c00449ba12ef4eb)

This is a little tool written in Go to generate the artifacts needed to run the
"1024 Microservices Test".

The objective of this test is to analyze and learn how a kubernetes cluster
behaves under a big load, like, for example, when deploying 1024 pods with 1024
services.

The tool will generate a number of kubernetes `deployment` artifacts based on the template located at `/tmpl`


## Usage

    ./1024apis -n=NUMBER_OF_DEPLOYMENTS -m=NUMBER_OF_DEPENDENCIES -d=TARGET_DIRECTORY

For example

    ./1024apis -n=500 -m=8 -d=my_test

This command will create 500 deployment files in the `my_test` directory. Each deployment file will launch
a container that can have __up to__ 8 dependencies.


### Please Note
This application does not deploy kubernetes artifacts!
