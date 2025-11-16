package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuildCmd = &cobra.Command{
	Use:   "codebuild",
	Short: "CodeBuild\n\nCodeBuild is a fully managed build service in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuildCmd).Standalone()

	})
	rootCmd.AddCommand(codebuildCmd)
}
