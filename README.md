[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/2sZOX9xt)
<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->




[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Best-README-Template</h3>

  <p align="center">
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template">View Demo</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Report Bug</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Request Feature</a>
  </p>
</div>
### Do not run this commands using fish shell

- [Contact](#Contact)
- [Built With](#built-with)
- [Features implemented](#features-implemented)
- [System Architecture](#system-architecture)
- [Prerequisites & Installation](#prerequisites--installation)
- [Starting Log Ingestion Service](#starting-log-ingestion-service)
- [Starting Query CLI Interface](#starting-query-cli-interface)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [License](#license)
- [Acknowledgments](#acknowledgments)



<!-- CONTACT -->
## Contact

#### Your Name - [@AiloniRohan](https://twitter.com/AiloniRohan)
#### Email:- rohanailoni@gmail.com
#### Additonal Email:- rohanailoni@gmail.com
#### phone:- 8008260370
#### Project Link: [https://github.com/dyte-submissions/november-2023-hiring-rohanailoni](https://github.com/dyte-submissions/november-2023-hiring-rohanailoni)

<p align="right">(<a href="#readme-top">back to top</a>)</p>




### Built With

* Programming Language:
  * GoLang
* External Services:-
  * RDS -- AWS
  * MongoDB--AWS
  * Redis  --Hosted By redis
* Frameworks:-
  * GIN - Powerful HTTP framework made over light weight threads goroutines.
  * Mongo - MongoDb HTTP driver
  * sqlgo  - Inbuild sql driver framework in go
  * cobra  - Powerful CLI framework in go


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Features implemented- All the below are implemented.
* Cli that can query all the fields--regular search ✔️
    * search by level
    * search by message
    * search by resourceId
    * search by timestamp
    * search by traceId
    * search by spanId
    * search by commit
    * search by metadata parentresourceid
* For fastness used prepared statements and effective sharding ✔️
* Distributed RDS(Managed by AWS) in different availability regions with read Replicas for faster read's--This is just a configuration of scaling group attached to RDS✔️
    * consistency and partition Tolerance is guaranteed by AWS across Read replica's ✔️
* Implement search within specific date ranges.    ---Bonus ✔️
* Implementing regular expression for Search         ---Bonus✔️
* Implementing allowing multiple filter                     ----Bonus✔️
    * Implemented only And Clause for multiple file like we can search only for `level` and `resourceId` ✅  but it is not supported to use `level` or `resourceId` ❌filter
* Provide real-time log ingestion and searching capabilities.we are using synchronous ingestion with SQL database so we are guaranteed with ACID(consistency ) ✅ so we can query real Time logs                                                                           -----Bonus ✔️
* Implement a simple Role Base search with authentication further explained in examples---Bonus ✔️
* Implemented wildcard search a shorted version of regex usually faster in sql DB's ---Additional Bonus ✔️
* Implemented Auth based filters via CLI ---BONUS ✔️

## System Architecture:-
### Synchronous Ingestion of Logs and Query CLI

![Image perssions removed](https://drive.google.com/uc?export=view&id=1xO_0KryYMbDjid--PKjltRFN5pAJ5pih)


## Overview:-

**Note:- All the Elaborate Reason for choosing shard Key,Database type(sql,Nosql),LoadBalancer,Query discussed below in HLD section**

#### Sample Log:-
```json
{
	"level": "error",
	"message": "Failed to connect to DB",
    "resourceId": "server-1234",
	"timestamp": "2023-09-15T08:00:00Z",
	"traceId": "abc-xyz-123",
    "spanId": "span-456",
    "commit": "5e5342f",
    "metadata": {
        "parentResourceId": "server-0987"
    }
}
```

#### Ingestion Service

* We are using Ngnix as the LoadBalancer with Configured Round Robin Algorithm and also used as reverse proxy so that we can scale our servers without changing the domain names.
* Now each Individual will use Database as according to shard key of level if it is  successful appropriate response is sent (we have leverage worker Pool with gorotuines with concurrent and parallel processing of request)
    * we also use go routines so that if there are some request concurrently running in the same server then a single MySQL Batch Request so the logs are Processed in single TCP request to server saving the bandwidth which significantly increases the performance.(Load test details are kept after **HLD** Below
* Also Leveraged the Connection Pool so that the server have enough set of connection set in `config.toml` so that first connection time is saved every TCP request and connection expiry(Normally connection expires after sometime but the connection pool notices and will start a new connection and add it to pool)

#### Query Service

* We offer 3 types of service for filtration type of all the field
    *  **Regular** Type search :- normal search in all rows and columns in the tables or with index if provided
    *  **Regex** Type search :- especially for searching regex values
    * **Wildcard** Type search :- especially for searching wildcard expression
        * Note;- wildcard expression are subset of regex but there are heavily optimized for sql queries so using wildcard instead of regex will have significant performance improvement
* Used Prepare statements method to leverage the fastness and secure of the query search with DB.


## Prerequisites && Installation

### System requirements:-

* Linux-Arch,Debian,Redhat
* GO-1.9
* Docker
* Docker-compose

### Services needed to run
* MYSQL   ------------------RDS HOSTED  IN AWS
    * Shard 1-----------------RDS HOSTED IN AWS
    * Shard 2 -----------------RDS HOSTED IN AWS
* Load Balancer(NGNIX) ---------defined config in dockerfile
* Log Ingestion Services ---------defined and configured docker-compose and dockerfile to handle have NGNIX as load balance and reverse proxy
* MongoDB (for authentication)                   ------------ Hosted in MongoDB Atlas

## Starting Log Ingestion Service

### Clone the Repo
```
git clone https://github.com/dyte-submissions/november-2023-hiring-rohanailoni.git
```

### Change the directory

```
cd server
```

### Run servers with Load Balancer with Docker-Compose

**Note:-** docker-compose should be installed and compatiable

```
docker-compose up
```

* This spins up ngnix and 4 Go-servers.
* Running the ngnix LoadBalancer running at  port 3000
* creates a connection pool to make request faster.


```


#### Successful Response
```json
{
	"status":"succesfully ingested the log"
}
```


## Starting Query CLI Interface
### Clone the Repo
```
git clone https://github.com/dyte-submissions/november-2023-hiring-rohanailoni.git
```

### Change the directory

```
cd server
```

### Install the dependencies
```
go mod download 
```


### Running the gocli
```
go run gocli.go --help
```

* This will return all the flags required for searching the database


<!-- USAGE EXAMPLES -->
### Usage Example with photos:-
* Find all logs with the level set to "error".
  * query cli :- `go run gocli.go --level "error"`
  
![use this endpoint drive.google.com/uc?export=view&id=123ovH4bs_4hK-NuaAnaISOVATMp-uoR4](https://drive.google.com/uc?export=view&id=123ovH4bs_4hK-NuaAnaISOVATMp-uoR4)

* Search for logs with the message containing the term "Failed to connect".
  * with just regular text search `go run gocli.go --message "Failed to connect"`
    * ![drive.google.com/uc?export=view&id=1PUfHok2piZi8xxNZLAhCcBrl5rGm8Wu8](https://drive.google.com/uc?export=view&id=1PUfHok2piZi8xxNZLAhCcBrl5rGm8Wu8)
  * with using regex text search `go run gocli.go regex --message "Failed to connect"`
    * ![https://drive.google.com/uc?export=view&id=14ibkxus3UIpB5Muf8SGJRtQssd3eG1mY](https://drive.google.com/uc?export=view&id=14ibkxus3UIpB5Muf8SGJRtQssd3eG1mY)
  * with using wildcard for text search `go run gocli.go wildcard --message "Failed to connect%"`
    * ![https://drive.google.com/uc?export=view&id=1grBMR943-ol9T67KhrrAOEm0Z6dJeyAm](https://drive.google.com/uc?export=view&id=1grBMR943-ol9T67KhrrAOEm0Z6dJeyAm)
* Retrieve all logs related to resourceId "server-1234"
  * `go run gocli.go --resourceId "server-1234"`
  * ![](https://drive.google.com/uc?export=view&id=1vVYlGLVug6ih2mV5PmVZz0J9ChMzyMbu)

* Filter logs between the timestamp "2023-09-10T00:00:00Z" and "2023-09-15T23:59:59Z"
  * query:-  `go run gocli.go --from "2023-09-10T00:00:00Z" --to "2023-09-15T23:59:59Z"`
  * can zoom and check timeline for more validity
    * ![check "https://drive.google.com/uc?export=view&id=1m__DBrFA6YjwIBsqzt2pIcM5SZOELG0B"](https://drive.google.com/uc?export=view&id=1m__DBrFA6YjwIBsqzt2pIcM5SZOELG0B)

## Usage
### Log Ingestor API Documentation

#### Ping Endpoint
```
GET 127.0.0.1:3000/ 
```

#### Response
```json
{
	"status":"Ping successfull"
}
```

#### Testing the Log Ingestor

Using post man for sending Post Request request
##### Log Ingestion Post endpoint
```http
POST 127.0.0.1:3000/
content-type "application/json"
{ "level": "error", "message": "Failed to connect to DB", "resourceId": "server-1234", "timestamp": "2023-09-15T08:00:00Z", "traceId": "abc-xyz-123", "spanId": "span-456", "commit": "5e5342f", "metadata": { "parentResourceId": "server-0987" }
```

``

#### Possible errors:
* Error Marshaling the request
* Error initializing the pool of connections to mysql
* Error reading the config file --mostly failed to ill configured TOML file
* Error Failed to prepare statements
* Error error executing SQL statements
### Query CLI Documentations


The Query Cli Has 3 functionalities
#### Authentication
```text
go run gocli.go auth -u <UserName> -p <Password>
```

This will authenticate the user and create a file in `HOMEDIR/.config/dyte/cli.json`
for every other request this will be used for auth


#### Regex Search
we use Mysql regex search functionality along with sql prepared statements this will query for most powerful regex engine provided by mysql

```
go run gocli.go regex --help
```

##### Output:-
```
Perform regex-based log filtering made for dyte

Usage:
  dyte regex [flags]

Flags:
  -p, --ParentResId string   Filter logs by Parent Resource ID pattern using regex
  -T, --TraceId string       Filter logs by Trace ID pattern using regex
  -c, --commit string        Filter logs by commit pattern using regex
  -h, --help                 help for regex
  -l, --level string         Filter logs by level pattern using regex
  -m, --message string       Filter logs by message pattern using regex
  -r, --resourceId string    Filter logs by resource ID pattern using regex
  -s, --spanId string        Filter logs by Span ID pattern using regex
```

#### Wildcard search ----Additional feature for faster search.
A wildcard character is used to substitute one or more characters in a string.

like passing `%rohan` will return all sentences ends with rohan, `rohan%` will return all sentences start with rohan, `%rohan%` will return all the sentences which has middle word `rohan`,
Many other patterns supported Read Here https://www.w3schools.com/sql/sql_wildcards.asp

```
go run gocli.go wildcard --help
```
##### Output:-
```text
Perform wilcardbased log filtering made for dyte

Usage:
  dyte wildcard [flags]

Flags:
  -p, --ParentResId string   Filter logs by Parent Resource ID pattern using wildcard(sql LIKE)
  -T, --TraceId string       Filter logs by Trace ID pattern using wildcard(sql LIKE)
  -c, --commit string        Filter logs by commit pattern using wildcard(sql LIKE)
  -h, --help                 help for wildcard
  -l, --level string         Filter logs by level pattern using wildcard(sql LIKE)
  -m, --message string       Filter logs by message pattern using wildcard(sql LIKE)
  -r, --resourceId string    Filter logs by resource ID pattern using wildcard(sql LIKE)
  -s, --spanId string        Filter logs by Span ID pattern using wildcard(sql LIKE)
```


#### Regular search
This is regular search used by MySQL leveraging prepared statements.

```
go run gocli.go --help
```

##### Output:-
```text
A simple log query processor made for dyte

Usage:
  dyte [flags]
  dyte [command]

Available Commands:
  auth        Perform auth for user registered for dyte .If not registerd ask admin
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  regex       Perform regex-based log filtering made for dyte
  wildcard    Perform wilcardbased log filtering made for dyte

Flags:
  -p, --ParentResId string   Filter logs by Parent Resource ID
  -T, --TraceId string       Filter logs by Trace ID
  -c, --commit string        Filter logs by commit
      --from string          Filter logs from timestamp
  -h, --help                 help for dyte
  -l, --level string         Filter logs by level
  -m, --message string       Filter logs by message
  -r, --resourceId string    Filter logs by resource ID
  -s, --spanId string        Filter logs by Span ID
  -t, --timestamp string     Filter logs from  this timestamp
      --to string            Filter logs to timestamp

Use "dyte [command] --help" for more information about a command
```

#### Using multiple fields with regex,regular,wildcard

```
go run gocli.go --level "error" regex --spanId "b[aeiou]bble" wildcard --message "%delectus"
```

The above will return all the logs with has level "error" and regex given spanId and wildcard match of messae.



##### Not supported ones(some Notes for CLI usage)
* cannot use a flag  multiple time like ex:-
```
	go run gocli.go --level "error" regex --level "error"
```

The above i am requesting level **Regular search** and for **Regex search** so we cannt use same field for all types of search
I have handled the error in code but this may break some deps so try not use it :)

* cannot have timestamp and to flag `to mean from which time it is required`
```
go run gocli.go --timestamp <timestamp> --to <timestamp>

```

this will result an error

* cannt have to time without from time mention but we can have from time flag without to flag as it will fallback to current time.

```
go run gocli.go --to <timestamp>  ---this is result an error
```

```
go run gocli.go --from <timestamp>  ---this is acceptable as it will fallback to current use timestamp.
```



## High Level Design
### Synchronous Log Ingestion and Query Cli 
![Image perssions removed](https://drive.google.com/uc?export=view&id=1xO_0KryYMbDjid--PKjltRFN5pAJ5pih)

#### Assumptions made for this Design as limited info was given.

1. The log is well-structured and may not evolve over time this is to make sure for using sql database and using sort keys for effective searching..
2. we get more queries on timestamp,resourceid,message search are the most searched queries so that we can indexed them faster.
3. This logs are very critical and should not be lost and this will reflect the current state of system
4. The logs we are getting from the upstream service where this is logged from is a distributed system
5. There are 2 possible values of `level` either `error` or `debug`; 


#### Round 1 of analysis of fields-My understanding
* `level` this is used to state the reason of logs this can be of type error,debug,info,warning for now we are choosing error,debug
* `message` This is used store the stage of the level like what is the message in log
* `resourceId` This is from the system we are getting the logs from a particular resource the system may evolve over time as it is considered a distributed system
* `timestamp` At which time this is logged.
* `Trace-Id` this is normal used to track the request in various services in a distributed system like `AWS-XRAY`
* `Span-ID` just a normal assumption of span
* `commit` This is a commit of the code for the log normally commit-ID
* `ParentResourceId` this is the name of the parent server from which this log is emitted  `ASSUMPTION-NO3`


## Choosing the ShardKey for Horizontal Sharding
This is an important step on what we chose on the key so if we look into the input format then the most frequent message that we get are
from either `level` or `resourceId` .

* Shardkey=`level` ✅
* Why are ignoring the `resourceId` ❌
  * we are thinking resourceId is some id of a system in distributed environment so the distributed environments are scalable with traffic so if we have different shards with resourceID then the sharding will evolve with increases the complexity 
  * but for level we know it is mostly fixed `debug` and `error` `Assumption 5`
  * 
## **Same Database, Different Tables:** or 2. **Different Databases:**
### 1. **Same Database, Different Tables:** ✅

#### Advantages:

1. **Simplicity:** It's often simpler to manage and query data within the same database.
2. **Easier Joins:** If you need to perform queries or reports that involve both error and info logs, having them in the same database makes JOIN operations straightforward.

#### Considerations:

1. **Table Size:** If the size of the combined table becomes very large, it might impact performance. However, modern databases can handle large tables efficiently.

### 2. **Different Databases:** ❌

#### Advantages:

1. **Isolation:** Having separate databases provides a level of isolation. Changes or issues in one database are less likely to impact the other.
2. **Scaling:** If you anticipate significantly different workloads for error and info logs, separating them into different databases can make it easier to scale and optimize each independently.

### Reasons to use first Type:-
* Simple to have the faster Joins in the SQL level which significantly decrease search time.


#### Considerations:

1. **Complexity:** Managing multiple databases can introduce complexity, especially when dealing with schema changes, backups, and migrations.
2. **Joins and Queries:** If you often need to query data that spans both error and info logs, having them in separate databases might complicate your queries.

## Choosing NoSql VS SQL Database

* From `Assumption 1` we can say our data is very well-structured so I am very inclined for SQL, so I am using AWS-RDS MYSQL instance
### Why Not NOSQL?
we are not talking about time series DB or key-value(dynamoDb)which is a suitable but is very costly , My opinion is on document store.
* The first requirement in the doc is that it is required for consistent i.e realtime search results i.e if a log is sent then it should be reflected in search Query and Nosql is eventually consistent.
* <u>First and foremost reason</u> is that when we are sharding joining of data should be handled by application level whereas we can leverage the SQL JOINS for Faster output and search time
* The  reason not to consider the NoSQL is increasing search time even with indexing as the engine have to look into every document for suitable env and 
* Searching on Indexed colum and regex is way faster in SQL than NoSQL.
* Disadvantages:-
  * NoSQL databases typically excel at key-value pairs or document storage but may have limited query capabilities compared to SQL databases, especially for complex queries and joins.


### Tradeoff by using SQL
* Compromising on faster write which is goto feature of NOSQL
* our schema will be flexible `Assumption 1` which might not be the case for growing system requirement always increases
* Inbuild sharding in NOSQL which is not the case for SQL we have to manually shard the DB


### Hybrid SQL AND NOSQL
* This is not used for time constraints
* we can use SQL for complex queries and Nosql for timeseries or some kind of analytical data
* The main problem implementing this is we are storing the log in 2 DB which might cause consistent issues and log may be missed as `Assumption 2` states logs are very important
* TO acheive this we have make Pub/Sub model then this will be asyncronus but problem statement require real time ..






## Indexing
* I am chosing `timestamp`,`ResourceId`,`message` which is mostly unique so i am choosing them as main Index.
* the decision to create indexes on "message," "resourceId," and "timestamp" depends on  specific use case and the queries your application executes. It's recommended to analyze your application's query patterns and perform testing to evaluate the impact of indexes on both read and write performance.
  ndexes depends on the types of queries your application performs most frequently. Indexes are used to speed up data retrieval, so it's generally a good idea to create indexes on columns that are frequently used in WHERE clauses, JOIN conditions, or ORDER BY clauses.

### I am using composite Index
```sql
CREATE TABLE ErrorLog (  
id SERIAL PRIMARY KEY,  
level VARCHAR(255),  
message TEXT,  
resourceId VARCHAR(255),  
timestamp TIMESTAMP,  
traceId VARCHAR(255),  
spanId VARCHAR(255),  
commit VARCHAR(255),  
metadata_parentResourceId VARCHAR(255)  
);



CREATE INDEX idx_log_timestamp_resourceId ON ErrorLog (timestamp,message,resourceId);

```

### TradeOFF using this Indexing

#### Our main question is on increasing the search efficeicy which is inturn acheived by indexing but there are some tradeoff

- **Disk Space:** Indexes consume additional disk space. Ensure that you have sufficient storage capacity.

- **Query Patterns:** Indexes should align with your application's query patterns. If you rarely filter or sort by a particular column, an index on that column might not be as beneficial.

- **Composite Indexes:** Depending on the types of queries, you might also consider composite indexes that cover multiple columns.

## Connection Pool
we are using connection pool for reducing the TCP on every connection using go library `sqlx`
* Connection pools in Go efficiently reuse established database connections, minimizing the overhead of connection creation and reducing latency in database operations
* Connection pooling enhances the scalability of Go applications by managing a fixed set of connections that can be shared among multiple goroutines, preventing resource exhaustion during high demand.
* Connection pools help manage system resources effectively by controlling the number of concurrent connections to the database, preventing issues such as resource exhaustion and optimizing resource utilization.
* Connection pools provide consistent performance by reusing connections, and they offer concurrency support, ensuring safe acquisition and release of connections across multiple goroutines.

## Authentication:-
I am using mongoDB authentication via cli to check the in DB and the sample document is the format of 

```json
{"_id":
{"$oid":"6558f317be1f48bfdd9147bb"},
  "user":"rohan",
  "password":"1234",
  "Access":{"regex":["all"],"regular":["all"],"wildcard":["all"]}}
```
now i can run all the filters in logs

but if i changed the above doc into 
```json
{"_id":
{"$oid":"6558f317be1f48bfdd9147bb"},
  "user":"rohan",
  "password":"1234",
  "Access":{"regex":["level"],"regular":["all"],"wildcard":["all"]}}
```
then i can just only go with level

```
go run gocli.go regex --message
```

this will return an error as this is not authrorised for the user real examples is attached in ## Usage with real example session


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>











<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
