# LocalCIJ

A simple watch-then-execute agent. I use it to deploy Cordova and Ionic apps to multiple devices.

## Quick Start

Have shell scripts you want to run ready in your source repository.

As an example, for 3 files called `deploy-ios`, `deploy-android` and
`deploy-web`, run:

```
$ localci deploy-ios deploy-android deploy-web
```

LocalCI is now watching those files. To activate a job, just `touch` the appropriate file.
A convenient setup is a shortcut via your favorite editor, or simply
through your terminal.



# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2015 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



