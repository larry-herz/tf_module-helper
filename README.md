# tf Module Creation Helper

## Purpose

This program will assist in the creation of terraform modules by creating the folder and a few basic files needed when creating a module.

## How to use

* Download the version for the platform where development will take place. There are versions for MacOS, Linux, and Windows on x86_64 architecture.
* Copy the binary to a directory in the PATH.
* Change directories to the directory where modules will be stored i.e. the `modules` directory.
* When creating a new module, from a command prompt / terminal window, execute:

  ```
  tf module <modulename>
  ```

  Where `<modulename>` is the name of the  module to be created.

## What this program does

This program does the following:

* Create a folder named `<modulename>` in the current directory.
* Create an empty `main.tf` file in the `<modulename>` directory.
* Create an empty `variables.tf` file in the `<modulename>` directory.
* Create an empty `outputs.tf` file in the `<modulename>` directory.
* Create a skeleton of a `README.md` file in the `<modulename>` directory.
* Create a .gitignore file with basic content in the `<modulename>` directory.
* Create a `template` folder in the `<modulename>` directory.
* Create a skeleton `<modulename>.tf` file in the `template` directory.
* Create a `AWS-providers.txt` file in the `template` directory, if the module name begins with AWS. (The contend of this file is needed for initial manual testing.)

## What this program does not do

Create any actual terraform content, beyond the minimal skeleton of code in the `<modulename>.tf` file.

## Caveats

This is very much alpha grade code and while it does what is described above, there is little error handling, and no input checking at this time. The code is built with the possibility of expansion in mind, but the first parameter passed on the command line currently does nothing. It is however required.