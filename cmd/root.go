package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)
}

var rootCmd = &cobra.Command{
	Use:   "pop {project [branch] | repositoryUrl}",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Open project or repository in directories",
	Run: func(cmd *cobra.Command, args []string) {
		parsedArgs, err := ParseFromParts(args)
		fmt.Println(parsedArgs)
		cobra.CheckErr(err)

		projectInDirectory, err := FindInDirectories(parsedArgs.Project, GetDirectories())
		cobra.CheckErr(err)

		if parsedArgs.HasBranch() {
			gitFetch := exec.Command("git", "-C", projectInDirectory, "fetch")
			cobra.CheckErr(gitFetch.Run())

			gitCheckout := exec.Command("git", "-C", projectInDirectory, "checkout", parsedArgs.Branch)
			cobra.CheckErr(gitCheckout.Run())
		}

		openProject := exec.Command(GetEditor(), projectInDirectory)
		openProject.Stdin = os.Stdin
		openProject.Stdout = os.Stdout
		openProject.Stderr = os.Stderr
		cobra.CheckErr(openProject.Run())
	},
}

type Args struct {
	Project string
	Branch  string
}

func (args Args) HasBranch() bool {
	return args.Branch != ""
}

func ParseFromParts(parts []string) (Args, error) {
	projectOrRepository := parts[0]

	var project = ""
	var branch = ""

	if !strings.Contains(projectOrRepository, "/") {
		project = projectOrRepository
	}

	parsed, _ := url.Parse(projectOrRepository)
	splitPathWithEmpties := strings.Split(parsed.Path, "/")
	var splitPath []string
	for i := range splitPathWithEmpties {
		if splitPathWithEmpties[i] != "" {
			splitPath = append(splitPath, splitPathWithEmpties[i])
		}
	}

	if len(splitPath) >= 2 {
		project = splitPath[1]
	}

	if len(splitPath) == 4 {
		branch = strings.Join(splitPath[3:], "/")
	}

	if len(parts) == 2 {
		branch = parts[1]
	}

	if project == "" {
		return Args{}, errors.New("Could not parse project project")
	}

	return Args{project, branch}, nil
}
