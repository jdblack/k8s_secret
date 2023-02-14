
# Overview

This program decodes and displays the strings within a K8s secret becuase `-o
jsonpath='{.data.*}' | base64 -d`  gets annoying.   


# building
`  go build .; go install . `

# Usage

Just run  `secret -ns NAMESPACE -s SECRET` and get your secrets.  Namespace is
optional, of course

```
~$ secret -ns kube-auth -s ldap
Secret                      Value
LDAP_ADMIN_PASSWORD         nauhngehssabel 
LDAP_CONFIG_ADMIN_PASSWORD  nauhngehssabel
```
