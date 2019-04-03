# awscaller
CLI for shows aws api caller infomations.

# Installation
```
$ git clone https://github.com/tabakazu/awscaller
$ cd awscaller
$ go build
```

# Usage
```
$ ./awscaller
- Account  : 000000000000
- UserId   : ABCDEFGHIJKLMNOPQRSTU
- UserName : tabakazu
- AttachedPolicies :
    - AdministratorAccess
```

## Commands
```
$ ./awscaller account
- Account : 000000000000

$ ./awscaller userid
- UserId : ABCDEFGHIJKLMNOPQRSTU

$ ./awscaller username
- UserName : tabakazu

$ ./awscaller policies
- AttachedPolicies :
    - AdministratorAccess
```