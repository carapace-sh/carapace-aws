package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Gets a list of build project names, with each build project name representing a single build project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listProjectsCmd).Standalone()

	codebuild_listProjectsCmd.Flags().String("next-token", "", "During a previous call, if there are more than 100 items in the list, only the first 100 items are returned, along with a unique string called a *nextToken*.")
	codebuild_listProjectsCmd.Flags().String("sort-by", "", "The criterion to be used to list build project names.")
	codebuild_listProjectsCmd.Flags().String("sort-order", "", "The order in which to list build projects.")
	codebuildCmd.AddCommand(codebuild_listProjectsCmd)
}
