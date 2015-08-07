env-template
============

Very simple command that reads a file, replaces placeholders with the contents of environment
variables, and outputs to stdout.

This can help with, for example, creating a configuration file for use within Docker.

Install
-------

For Linux x86_64, you can grab a release and place it anywhere, for example:

    sudo curl -L -o /usr/local/bin/env-template \
        https://github.com/fazy/env-template/releases/download/0.1/env-template-0.1-linux-x86_64 \
    && sudo chmod 0755 /usr/local/bin/env-template

For other systems you can clone this repository, [download Go](https://golang.org/doc/install)
and [build with one command](https://gobyexample.com/hello-world).

Usage
-----

This program takes one argument, the name of your template file, and always outputs the
result to stdout:

    env-template <templatefile>

For example:

    env-template config.yml.tpl > config.yml

Template file
-------------

The template file is based on [Go Template](http://golang.org/pkg/text/template/), with
all the variables from your external environment already set.

Here's a simple exmaple:

**test.yml**

    User: {{ env "USER" }}
    Path: {{ env "PATH" }}
