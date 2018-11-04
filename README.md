# yaml-printer

a tool to print yaml file

go build

Then add `yaml-printer` binaryFile to your `PATH`.

cat a-yaml-file.yaml | yaml-printer

> example:

```
$ cat configtx.yaml | yaml-printer 
Application:
  Capabilities:
    V1_2: true
  Organizations: null
  Policies:
    Admins:
      Rule: MAJORITY Admins
      Type: ImplicitMeta
    Readers:
      Rule: ANY Readers
      Type: ImplicitMeta
    Writers:
      Rule: ANY Writers
      Type: ImplicitMeta
Capabilities:
```