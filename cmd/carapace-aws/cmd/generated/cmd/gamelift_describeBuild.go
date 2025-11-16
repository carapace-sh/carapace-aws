package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeBuildCmd = &cobra.Command{
	Use:   "describe-build",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves properties for a custom game build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeBuildCmd).Standalone()

		gamelift_describeBuildCmd.Flags().String("build-id", "", "A unique identifier for the build to retrieve properties for.")
		gamelift_describeBuildCmd.MarkFlagRequired("build-id")
	})
	gameliftCmd.AddCommand(gamelift_describeBuildCmd)
}
