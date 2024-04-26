# Auto upload

Auto upload is a program that auto upload all files in a specified directory into min.io.

## Installation:

Create a .env file and fill it with correct information (the one here are good for local development only)

````
MINIO_ENDPOINT=127.0.0.1:9000 #The API endpoint of your minio
MINIO_ACCESS_KEY=minioadmin #Username
MINIO_SECRET_ACCESS_KEY=minioadmin #Password
MINIO_BUCKET=uploads #The bucket the application will use (It is created by the application)

ROOT_DIR_PATH=/home/namu/Téléchargements/ #The directory you want to upload
````

Note: There MUST be a "/" at the end of the ROOT_DIR_PATH env variable.

Then just run the executable version.

## build

For people who don't find an executable for your OS, here is how to build the application.

First, you need Go, at least version 1.22.

Then clone this repo with

````
git clone https://github.com/Namularbre/fileUpload.git
````

Install all the dependencies with

````
go mod download
````

And run

````
go build 
````

There is no real "Release mode" for go, we can omit some flags, but it will remove interesting data if the program crashes.
