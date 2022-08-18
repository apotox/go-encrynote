## [Encrynote Website](https://encrynote.safidev.de) 

**Encrynote** – an application Built on AWS Lambda and other next-gen cloud services, that auto-scale and only charge you when they run. This is a working example about how to lowers the total cost of running and operating your apps, enabling you to build more and manage less.

Encrynote is a web application tool that uses serverless framework to deploy both react frontend and cloud infrastructure needed to make the serverless **Go** backend works. using Reactjs, Typescript, Go, Mongodb, Github actions, and SQS service.
you can find database migrations, unit tests , e2e tests and how to mock aws services, also check the github actions to find how to use test and deploy Docker containers with custom github action.

the purpose of this Application is to encrypt notes (messages) and share a one-time link with another party, i used aes for encryption. you cant use this app for production its just a POC.

## Three functions:
### in /functions
- create_note function // called from API gateway by a POST request to create new Note
- read_note function // called from API gateway by a GET request to read a note by its id and secretKey
- delete_note function // triggred by SQS queue (QueueDeleteNote), the read_function publishs to the queue a NoteId after reading the note. delete_note function will consume this id and delete the Note.

each function will get compiled to a **Go** binary and deployed as a lambda function (see Makefile)



Actively maintained by [Safi](https://dz.linkedin.com/in/safi-eddine-bouhentala).