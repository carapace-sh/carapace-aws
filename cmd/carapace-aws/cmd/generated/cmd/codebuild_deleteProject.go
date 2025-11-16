package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes a build project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteProjectCmd).Standalone()

	codebuild_deleteProjectCmd.Flags().String("name", "", "The name of the build project.")
	codebuild_deleteProjectCmd.MarkFlagRequired("name")
	codebuildCmd.AddCommand(codebuild_deleteProjectCmd)
}
