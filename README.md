# DERgo
---

![Static Badge](https://img.shields.io/badge/License-MIT-green)
![Static Badge](https://img.shields.io/badge/stability-experimental-red)

Does everything resolve (go version)
---

## What is DERgo?
This is a __toy project__ that is intended mostly as a way to practice coding in golang, but is inspired by a real world
use case I have of wanting a tool that can:
- Be easily portable and have minimal external dependencies in order to run.
- Alter expectations based on environmental considerations within the a common configuration.
- Help identify and validate DNS is still resolving correctly after making configuration changes to one or more DNS servers.
- Use a configuration file, that can be deployed and read by the tool in a variety of situations to ensure even in complex DNS hiererarchies clients can resolve as expected after making changes.
---

## Features
- Takes a yaml file (format described below) with a list of hosts and optionally expectations, and checks that the hosts resolve using DNS.
```yaml
---
records:
  - name: kernel.org # If kernel.org resolves to anything this will be considered a pass
  - name: example.com # If 10.10.10.10 is not in the list of records returned when resolving example.com this will fail
    expect: 93.184.215.14
  - name: someplace.org # If DERgo is run with an environment specified other that dev or prod this will not fail regardless of if it resolves to the expected address
    expect: 10.10.10.10
    environments:
      - dev
      - prod
```
---

## TODO
- Read from a file or take command line input
- Make work with reverse lookups
- Make work with TXT records
- Use the local resolver, or a specified name server
- Implement a way for dynamic record validations (e.g. hosts behind cloudflare proxies)
