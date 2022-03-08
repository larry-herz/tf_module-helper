package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"strings"
)

var modulename = os.Args[2]
var cloud string

func main() {

	_, err := os.Stat(modulename)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(modulename, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	os.Chdir(modulename)

	createFile("main.tf")
	createFile("variables.tf")
	createFile("terraform.tfvars")
	createFile("outputs.tf")
	createFile(".gitignore")
	writeFile(".gitignore")
	createFile("README.md")
	writeFile("README.md")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("template", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	os.Chdir("template")

	createFile(fmt.Sprintf("%s.tf", modulename))
	writeFile(fmt.Sprintf("%s.tf", modulename))

	// Check for the AWS prefix on the module, if found create the AWS-providers.txt file for testing
	cloud = strings.Split(modulename, "-")[0]
	switch cloud {
	case "aws":
		createFile("AWS-providers.txt")
		writeFile("AWS-providers.txt")
	}
}

func createFile(filename string) {
	// check if file exists
	var _, err = os.Stat(filename)

	// create file if not exists
	if os.IsNotExist(err) {
			var file, err = os.Create(filename)
			if isError(err) {
					return
			}
			defer file.Close()
	}
	fmt.Println("File Created Successfully", filename)
}
func writeFile(filename string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	switch filename {
	case "README.md":

		// Write some text line-by-line to file.
		_, err = file.WriteString(fmt.Sprintf("# %s\n\n", modulename))
		if isError(err) {
				return
		}
		_, err = file.WriteString("## This module builds (what is the module for)\n\n")
		if isError(err) {
				return
		}

		_, err = file.WriteString("### The following variables are necessary to run this module\n\n* name - (Required) Specifies the name of the resource. Changing this forces a new resource to be created.\n\n")
		if isError(err) {
				return
		}

		_, err = file.WriteString("### To run this module from a root module\n\n")
		if isError(err) {
			return
		}

		_, err = file.WriteString(fmt.Sprintf("```terraform\nmodule \"%s\" {\n  source = \"git::git@ssh.dev.azure.com:v3/thinkahead-azure/client-scadm/%s?ref=master\"\n  tags = {\n    \"key\" = \"value\"\n  }\n}\n```\n", modulename, modulename))
		if isError(err) {
			return
		}
	case fmt.Sprintf("%s.tf", modulename):
		// Write some text line-by-line to file.
		_, err = file.WriteString("### FOR MANUAL TESTING - Copy the contents of the aws-providers repo into this space. ###\n\n")
		if isError(err) {
				return
		}

		_, err = file.WriteString(fmt.Sprintf("module %s {\n\n", modulename ))
		if isError(err) {
				return
		}

		_, err = file.WriteString(fmt.Sprintf("  source = \"git::git@ssh.dev.azure.com:v3/thinkahead-azure/client-scadm/%s?ref=master\"\n }", modulename))
		if isError(err) {
			return
		}

	case ".gitignore":
		_, err = file.WriteString("# .tfstate files\n*.tfstate\n*.tfstate.*\n\n# Crash log files\ncrash.log\n\n")
		if isError(err) {
				return
		}
		_, err = file.WriteString("#ignore any data contained in the .terraform directory\n**/.terraform\n\n# Ignore .tfvars files generated for a terraform run.\n**/terraform.tfvars\n**/testing.tfvars\n")
		if isError(err) {
				return
		}
		_, err = file.WriteString("# Ignore any macOS related extra window size and position data.\n**/.DS_Store\n\n**/aws-providers.txt\n\n**/test.txt")
		if isError(err) {
				return
		}

	case "AWS-providers.txt":
		_, err = file.WriteString("terraform {\n  required_providers {\n    aws = {\n      source = \"hashicorp/aws\"\n      version = \"~>3.1.0\"\n    }\n    random = {\n      source=\"hashicorp/random\"\n      version = \"~>2.3.0\"\n    }\n  }\n  required_version = \"0.13\"\n}\n\n")
		if isError(err) {
				return
		}

		_, err = file.WriteString("provider \"aws\" {\n  profile = var.profile \n  region = var.region\n\n  assume_role { \n    role_arn = var.arn_name\n    external_id = var.ext.id\n  }\n}\n")

		if isError(err) {
				return
		}
	}
		// Save file changes.
		err = file.Sync()
		if isError(err) {
				return
		}
	fmt.Println("File Updated Successfully.")

}

func readFile(filename string) {
	// Open file for reading.
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if isError(err) {
			return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
			_, err = file.Read(text)

			// Break if finally arrived at end of file
			if err == io.EOF {
					break
			}

			// Break if error occured
			if err != nil && err != io.EOF {
					isError(err)
					break
			}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func deleteFile(filename string) {
	// delete file
	var err = os.Remove(filename)
	if isError(err) {
			return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
			fmt.Println(err.Error())
	}

	return (err != nil)
}
