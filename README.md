# commercial-shop

Commercial shop

## Why?

**Important part**: I need a nice and friendly way to deploy website server? But where are
solution?

### Goals?

**Important part:** A web server can serve basic things for commercial purpose.

**NOTE:** A sysadmin manage a web server is not a n00b. we will talk later
about friendly stuff.

**Extends goals:**
+ Secure web server.
+ Simple and easy to deploy web server?
+ Manage easy?
+ Other commercial database?

I will tell you later after we have achieved something.

Something in this documentation may not true but we will improve over time.

## Technology part?

### Database

+ Postgresql
  + Binary format
  + It is SQL

### Web API? (Backend)

+ Go
  + Fast (compiled language, Want to work like scripting? It can too.)
  + Memory safe (Garbage collector. Don't ask me why have you ever write C/C++?)
  + Data type
  + Design to be simple?
    + Force us to write the same code?
    + No duplicate loop syntax. `for` is love `for` is live
    + No class but provide struct and methods
    + No Override? Just copy paste and change like you want.
  + Package Management?
    + Unlike C/C++ where we try to pretend it have and packages depend on which
    platforms.
    + Not perfect but fit solution (Because Go packages use links which is bad
    idea)
  + Weird but fun? 

Can't use Go. Skill issue

Btw, we are using [Gin](https://github.com/gin-gonic/gin)

### Client (Frontend)

#### Web?

React?

We will talk about that later

#### Mobile?

We will talk about that later

### Server

Just use NixOS because it fit all the problems he have.

But why?

+ Need DevOps environment? NixOS
+ Docker? Nix.
+ An operating system? NixOS which is built on top of Nix package manager
+ Ansible? Nix files and NixOS for config and manage easily
+ Btrfs? ZFS with NixOS?
+ Backup way? NixOS can help you to roll back previous config

Simple part: Learn Nix

Hard part: Learn Nix

Bad part: You don't like to learn new technology and you realize that
computer sucks. But it is true fact. Stop using Computer we have been lied
ourself for thounsand decade.

Summary: I use these because it is fit not fun.

## Sad part

Only me.

