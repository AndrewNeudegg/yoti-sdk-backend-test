# Yoti SDK Back-end test

## Contents
- [Yoti SDK Back-end test](#yoti-sdk-back-end-test)
  - [Contents](#contents)
  - [Solution](#solution)
    - [Running the application](#running-the-application)
    - [Making a request](#making-a-request)
      - [Connectivity test](#connectivity-test)
      - [Service](#service)
    - [About](#about)
  - [Brief: Yoti SDK Back-end test](#brief-yoti-sdk-back-end-test)
    - [Introduction](#introduction)
    - [Goal](#goal)
    - [Input](#input)
    - [Output](#output)
    - [Deliverable](#deliverable)
    - [Evaluation](#evaluation)

## Solution
The solution is a Go application that serves a RESTful endpoint that accepts POST requests responding with a processed response.
POST requests Inputs and Outputs are stored within a local database.

### Running the application
To run the application ensure you have a valid Golang environment (Go >= 1.10), and have gathered any dependencies
using [dep](https://github.com/golang/dep), `dep ensure`. To keep Repo clean I will not include the `./vendor` directory
in the repository.

When the application is first run it will generate a configuration file `./config.json` and database file
`./application.db`. The contents of the configuration file are loaded once at application start. The contents are:

```json
{
    "log-configuration": {
        "describe-caller": true
    },
    "server-configuration": {
        "host": "localhost",
        "port": 8080
    },
    "database-configuration": {
        "db-name": "application.db",
        "lock-timeout": 1
    }
}
```

If you are running a service on port `8080` you may wish to change the port here: `"port": 8080`.

### Making a request

#### Connectivity test
To test connectivity:

```sh
curl localhost:8080
```

Should respond with:

```sh
Hello World!
```

#### Service
All service requests are provided by a POST to the `/` route, for instance the following command:

```sh
curl -d '{"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":"NNESEESWNWW"}' -H "Content-Type: application/json" -X POST http://localhost:8080/
```

Should return:

```sh
{"coords":[1,3],"patches":1}
```

### About

- The HTTP Router service is provided by the [mux](https://github.com/gorilla/mux) library.
- The logger is wrapped, but provided by [logrus](https://github.com/sirupsen/logrus) library.
- Database access is provided by [boltdb](https://github.com/boltdb/bolt). Project is archived, it is used here for simplicity.
- Whitebox and Blackbox tests are provided for various methods.
- Blackbox tests are run with a virtual copy of the application.

If an error occurs the logs are the best place to start, for instance:

```log
INFO[2019-05-20T16:20:37+01:00] [config.go:27:logging.ApplyConfiguration] Applied configuration to logging package.
INFO[2019-05-20T16:20:37+01:00] [config.go:18:server.ApplyConfiguration] Applied configuration to server package.
INFO[2019-05-20T16:20:37+01:00] [config.go:23:database.ApplyConfiguration] Applied configuration to database package.
INFO[2019-05-20T16:20:37+01:00] [database.go:20:database.connectDatabase] Attempting to connect to database.
INFO[2019-05-20T16:20:37+01:00] [database.go:24:database.connectDatabase] Connection made to database.
INFO[2019-05-20T16:20:37+01:00] [config.go:40:common.ApplyConfiguration] Loaded Configuration: &{LogConfig:{DescribeCaller:true} ServerConfig:{Host:localhost Port:8081} DatabaseConfig:{DBName:application.db LockTimeout:1}} .
INFO[2019-05-20T16:20:37+01:00] [routes.go:31:server.NewRouter] Added new route: POST:/
INFO[2019-05-20T16:20:37+01:00] [routes.go:31:server.NewRouter] Added new route: GET:/
INFO[2019-05-20T16:20:37+01:00] [server.go:67:server.(*SimpleServer).Start.func1] Server is starting on: http://localhost:8081
^CINFO[2019-05-20T16:20:38+01:00] [main.go:37:main.main.func1] Recieved OS signal: interrupt
INFO[2019-05-20T16:20:38+01:00] [tidy.go:10:common.Tidy] Requesting that packages tidy themselves.
INFO[2019-05-20T16:20:38+01:00] [utilities.go:7:database.Tidy] Tidying database package.
INFO[2019-05-20T16:20:38+01:00] [database.go:31:database.releaseDatabase] Attempting to release database.
INFO[2019-05-20T16:20:38+01:00] [utilities.go:9:database.Tidy] Done tidying database package.
INFO[2019-05-20T16:20:38+01:00] [main.go:39:main.main.func1] The application is now exiting.
```


## Brief: Yoti SDK Back-end test
### Introduction
You will write a service that navigates a imaginary robotic hoover (much like a Roomba) through an equally imaginary room based on:

- room dimensions as X and Y coordinates, identifying the top right corner of the room rectangle. This room is divided up in a grid based on these dimensions; a room that has dimensions X: 5 and Y: 5 has 5 columns and 5 rows, so 25 possible hoover positions. The bottom left corner is the point of origin for our coordinate system, so as the room contains all coordinates its bottom left corner is defined by X: 0 and Y: 0.
- locations of patches of dirt, also defined by X and Y coordinates identifying the bottom left corner of those grid positions.
- an initial hoover position (X and Y coordinates like patches of dirt)
- driving instructions (as cardinal directions where e.g. N and E mean "go north" and "go east" respectively)

The room will be rectangular, has no obstacles (except the room walls), no doors and all locations in the room will be clean (hoovering has no effect) except for the locations of the patches of dirt presented in the program input.

Placing the hoover on a patch of dirt ("hoovering") removes the patch of dirt so that patch is then clean for the remainder of the program run. The hoover is always on - there is no need to enable it.

Driving into a wall has no effect (the robot skids in place).

### Goal
The goal of the service is to take the room dimensions, the locations of the dirt patches, the hoover location and the driving instructions as input and to then output the following:

- The final hoover position (X, Y)
- The number of patches of dirt the robot cleaned up

The service must persist every input and output to a database.

### Input
Program input will be received in a json payload with the format described here.

Example:

```javascript
{
  "roomSize" : [5, 5],
  "coords" : [1, 2],
  "patches" : [
    [1, 0],
    [2, 2],
    [2, 3]
  ],
  "instructions" : "NNESEESWNWW"
}
```

### Output
Service output should be returned as a json payload.

Example (matching the input above):

```javascript
{
  "coords" : [1, 3],
  "patches" : 1
}
```

Where `coords` are the final coordinates of the hoover and patches is the number of cleaned patches.

### Deliverable
The service:

- is a web service
- must run on Mac OS X or Linux (x86-64)
- must be written in any of the languages that we support with our SDKs (Java, C#, Python, Ruby, PHP, Node, Go)
- can make use of any existing open source libraries that don't directly address the problem statement (use your best judgement).

Send us:

- The full source code, including any code written which is not part of the normal program run (scripts, tests)
- Clear instructions on how to obtain and run the program
- Please provide any deliverables and instructions using a public Github (or similar) Repository as several people will need to inspect the solution

### Evaluation
The point of the exercise is for us to see some of the code you wrote (and should be proud of).

We will especially consider:

- Code organisation
- Quality
- Readability
- Actually solving the problem

This test is based on the following gist https://gist.github.com/alirussell/9a519e07128b7eafcb50
