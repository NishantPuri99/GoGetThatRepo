# GoGetThatRepo
### Nishant Puri
---
GoGetThatRepo (pun intended), is a Go Project aimed at doing the following:
Build two Golang REST APIs that,
1. Firstly perform a **SCAN operation (POST /scan)**,
   1. Fetch JSON objects from a remote [GitHub repository](https://github.com/velancio/vulnerability_scans)
   2. Performing the Fetch concurrently with retries (3 max)
   3. Parse the received JSON objects to store them in a local sqlite db instance (using gorm)
2. Secondly perform a **QUERY operation (POST /query)**,
   1. Basis a filter value, (currently just 'severity'), fetch the Critical Vulnerabilities stored in the sqlite db and return them to the user

**Build and Run**
1. Clone the repo
```gh repo clone NishantPuri99/GoGetThatRepo```
2. Navigate to the directory created do a docker build ```cd GoGetThatRepo``` and then ```docker build -t github-scanner . ```
3. Once the build is complete run the docker container with ```docker run -p 8080:8080 github-scanner```
4. A message with ```Database and tables created successfully.``` pops up on the terminal

_**Open to a REST API Client (like Postman).**_
For the Scan API:
- Send a **POST** request to 'localhost:8080/scan'
- With payload as follows ```{
  "repo": "velancio/vulnerability_scans",
  "files": ["vulnscan15.json","vulnscan16.json"]
}```
- The files in the payload can vary as per the files in the repo

_Once the JSONs are parsed and added to the SQLdb then for the Query API_
- Send a **POST** request to 'localhost:8080/query'
- With payload as follows ```{
    "filters": {
        "severity": "CRITICAL"
    }
}```
- The Severity can vary from **CRITICAL**, **HIGH**, **MEDIUM** and **LOW**


✅ **Completed work**
1. The files fetched from the GitHub repo are done parallelly using goroutines
2. The code is robust as far as manual testing has been performed and handles success and failure cases
3. the code is modular and organized into different folders
4. Easily built on a docker container using the DockerFile provided

⏳ **Pending work**
1. This application has not been unit-tested due to time constraints.
2. The concurrency part would require performance testing.
