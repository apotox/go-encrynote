## [Encrynote Website](https://encrynote.safidev.de) 


![arch](https://github.com/apotox/go-encrynote/blob/master/client/app/public/screen.png?raw=true)

**Encrynote** – an application Built on AWS Lambda and other next-gen cloud services, that auto-scale and only charge you when they run. This is a working example about how to lowers the total cost of running and operating your apps, enabling you to build more and manage less.

Encrynote is a web application tool that uses serverless framework to deploy both react frontend and cloud infrastructure needed to make the serverless **Go** backend works. using Reactjs, Typescript, Go, Mongodb, Github actions, and SQS service.
you can find database migrations, unit tests , e2e tests and how to mock aws services, also check the github actions to find how to use test and deploy Docker containers with custom github action.

the purpose of this Application is to encrypt notes (messages) and share a one-time link with another party, i used aes for encryption. you cant use this app for production its just a POC.

## AWS Lambda Functions:
### in `/functions`
- create_note function // called from API gateway by a POST request to create new Note
- read_note function // called from API gateway by a GET request to read a note by its id and secretKey
- delete_note function // triggred by SQS queue (QueueDeleteNote), the read_function publishs to the queue a NoteId after reading the note. delete_note function will consume this id and delete the Note.

each function will get compiled to a **Go** binary and deployed as a lambda function (see Makefile).

## SQS 
the `delete_function` triggred by an `SQS` message, which is published after the user read a note.
the message is the note's `_id`.

## Github Actions and CI/CD

### Testing:

every `function` will be tested separately using a common docker image `action_tester.Dockerfile` , and by usign github action Matrix we can test all function in parallel ,
you can add more functions to be tested in `.github/workflows/main.yml` 

``` yaml
test_functions_job:
    runs-on: ubuntu-latest
    name: Test functions
    strategy:
      fail-fast: true
      matrix:
        function:
          [
            # { functionName: "<function package name>" },
            { functionName: "create_note" },
            { functionName: "read_note" },
            { functionName: "delete_note" },
          ]
```

the custom github action `.github/actions/function-test` will be responsible for creating/running the test docker containers.

## Deployment:

after a successful tests, a deploy job will be started to deploy all the compiled lambda functions. i used `serverless framework v3` to make that easier. you can see the `serverless.yaml`.

the Deploy action uses `action_deployer.Dockerfile` to:

  - build the `Go` functions
  - install the necessary `Nodejs` dependencies to make the serverless framework works.
  - create a production config json file (decoding base64 secrets to a new file config.prod.json) this file used by the `serverless cli` to set all the required envirenment var during the deploy-to-aws process.

the `entrypoint shell scripts` used to handle the diffrent args from the custom github actions.

## Database
- Mongodb database

## Migrations
- you can add new migrations to `./migrations`
```go
// ...
migrate.Register(func(db *mongo.Database) error {
		err := db.CreateCollection(context.Background(), "<collection name>", nil)
		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		fmt.Printf("Dropping %s collection\n", "notes")
		err := db.Collection("<collection name>").Drop(context.TODO())
		if err != nil {
			return err
		}
		return nil
	})
```

## Tests
beside the `unit tests` , there are the e2e tests to test the function behavior with mocked requests/events.


maintained by [Safi](https://dz.linkedin.com/in/safi-eddine-bouhentala).