package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_stopBuildCmd = &cobra.Command{
	Use:   "stop-build",
	Short: "Attempts to stop running a build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_stopBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_stopBuildCmd).Standalone()

		codebuild_stopBuildCmd.Flags().String("id", "", "The ID of the build.")
		codebuild_stopBuildCmd.MarkFlagRequired("id")
	})
	codebuildCmd.AddCommand(codebuild_stopBuildCmd)
}
