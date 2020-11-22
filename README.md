# Motivation
Once upon a time, my internet was sporadically going out and my ISP was of no real help,
thinking that just resetting the modem would fix the issue. I created this lightweight
tool to keep an audit of when the internet was connected and disconnected to show them
that it was a real issue, and to try to identify any patterns in the up/downtime.

# Installation
```shell script
git clone https://github.com/danielfullerton/wifi-connection-analyzer.git
```

# Usage
### Quickstart
The quickstart script enables you to build and run the analyzer daemon with
a single shell script:
```shell script
./scripts/quickstart.sh
```
This will output logs at **~/analyze.csv** and will run every 60 seconds.

### Running with Custom Arguments
First, build the project:
```shell script
go build cmd/main.go
```

Next, run the outputted binary:
```shell script
./main -i 30 -h "http://clients3.google.com/generate_204" ~/custom_analyze_log.csv
```

This will start the daemon script in the background.
### Argument List
> - **-h** : The URL you wish to use as a status check endpoint (optional).
> - **-i** : The interval, in seconds, at wish you want the daemon to check your connection (optional).

The final argument of the tool is the path to the file that you wish to use for your CSV logs to be written to.

### Killing the Daemon
```
./scripts/kill.sh
```
This will kill the PID that contains the running daemon. Note that if you start up multiple instances
of the daemon, this script will only kill the most recent process. You will need to manually find
and kill the other PIDs.

# Contributors
[Daniel Fullerton](https://github.com/danielfullerton)

# License
The MIT License (MIT)

Copyright (c) 2020 Daniel Fullerton

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
