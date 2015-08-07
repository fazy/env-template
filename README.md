env-template
============

Very simple command that reads a file, replaces placeholders with the contents of environment
variables, and outputs to stdout.

This can help with, for example, creating a configuration file for use within Docker.

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
