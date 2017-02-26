# thunks

`thunks` wraps a single executable file in a self-extracting JAR file, to be used in environments that expect executable JARs. 
The command line passed to the JAR file will be passed to the underlying executable.

## Usage

```sh
$ thunks /bin/ls
$ java -jar ls.jar /
acct  bin  boot  cache  data  dev  etc  home  init  lib  lib64  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```

## How does it work?

Thunks creates a JAR containing a manifest, plug two different files:
- An executable file
- A Java class

The Java class will extract the executable from the archive to a temporary directory, forwarding to it the command line arguments originally passed to it.
