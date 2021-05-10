# Robot Warehouse

## Scenario

We are installing a new type of robot into our (hypothetical) warehouse as part of an automation project.  As part of this project, there are various software components which need to be developed.

## About the Robots

For convenience the robot moves along a grid in the roof of the warehouse and we have made sure that all of our warehouses are built so that the dimensions of the grid are 10 by 10; objects in the warehouse, including the robots, are always aligned with the grid, so object locations' may be treated as integer coordinates.  We've also made sure that all our warehouses are aligned along north-south and east-west axes. The system operates on a cartesian coordinate map that aligns to the warehouse's physical dimensions: point (0, 0) indicates the most south-west and (10, 10) indicates the most north-east.

Each robot operates by being given 'tasks' which each consist of a string of 'commands':

All of the commands to the robot consist of a single capital letter and different commands are optionally delineated by whitespace.

The robot should accept the following commands:

- N move one unit north
- W move one unit west
- E move one unit east
- S move one unit south

Example command sequences:

* The command sequence: `"N E S W"` will move the robot in a full square, returning it to where it started.

* If the robot starts in the south-west corner of the warehouse then the following commands will move it to the middle of the warehouse: `"N E N E N E N E"`

The robot will only perform a single task at a time: if additional tasks are given to the robot while is busy performing a task, those additional tasks are queued up, and will be executed once the preceding task is completed (or aborted for some reason).  Each task is identified with a unique string ID, and a task which is either in progress or enqueued can be aborted/cancelled at any time.  If the robot is unable to execute a particular command (for instance, because the command would cause the robot to run into the edges of the warehouse grid) then an error occurs, and the entire task is aborted.

## How to Run API
```
cd spot-restfulapi
go run main.go
```
Running on `http://localhost:8081/`

## How to Run Test
```
cd spot-restfulapi
go test ./...
```

## Challenge
Question
The ground control station wants to be notified as soon as the command sequence completed. Please provide a high level design overview how you can achieve it. This overview is not expected to be hugely detailed but should clearly articulate the fundamental concept in your design.

### Architecture Design
![image](https://user-images.githubusercontent.com/5085888/117626097-d9c4de00-b1ca-11eb-933c-b5c2106ca536.png)

This is a high level architecture design of how a control station (i.e. web/mobile app) can be notified when a command sequence is completed. The following is a brief explanation of each component.
 
The API wwould be able to post command, delete command, and get current status via the robot's SDK as described in the requirement of the restful API. The API would also store the command (taskID, the command string, timestamp) in a database. This design is based on the robot being able to push sensory data into a data stream (e.g. AWS Kinesis). Data such as the location of the robot can be pushed to a service which would constantly check if it matches the final coordinate of the latest command. If coordinates show that the command has been completed a websocket API (AWS API Gateway) can push the completion status to a web/mobile app.

If the robot is unable to stream sensory data, a service would need to compute the completion status based on the robot's starting position, the command and current position.

source: https://drive.google.com/file/d/111tTA61ZPmBx6S8atg8FVMYYdAwMydUR/view?usp=sharing
