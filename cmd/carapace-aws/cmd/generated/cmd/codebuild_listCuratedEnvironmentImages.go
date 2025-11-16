package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listCuratedEnvironmentImagesCmd = &cobra.Command{
	Use:   "list-curated-environment-images",
	Short: "Gets information about Docker images that are managed by CodeBuild.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listCuratedEnvironmentImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listCuratedEnvironmentImagesCmd).Standalone()

	})
	codebuildCmd.AddCommand(codebuild_listCuratedEnvironmentImagesCmd)
}
