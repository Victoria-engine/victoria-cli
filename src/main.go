package main

import (
	"flag"
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// REPO : The repository url for the template
const REPO = "git clone https://github.com/Victoria-engine/blog-template"

func main() {
	pathname := setupArguments()

	if flag.NFlag() == 0 {
		printUsage()
	}

	if pathname != "" {
		fmt.Printf("Creating your new Victoria blog to: ./%v \n\n...\n\n", pathname)
		// Clone the given repository to the given directory

		_, err := git.PlainClone(pathname, false, &git.CloneOptions{
			URL:      "https://github.com/Victoria-engine/blog-template",
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Print how to use the project
		printInstructions(pathname)
	}
}

func setupArguments() string {
	var pathname = flag.String("new", "my-victoria-blog", "the name of the project")

	flag.Parse()

	return *pathname
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func printInstructions(pathname string) {
	fmt.Printf("\n\nWell done! Now you can: \n\n1. cd %v \n2. yarn develop or npm install to install the dependencies.\n", pathname)
	fmt.Println("3. And don't forget to create a .env file in the root of your project and add the REACT_APP_VICTORIA_KEY='your_key' !")
	fmt.Println("\n4. yarn start to run the project !\n")

	fmt.Println("Happy blogging !\n\n")
}
