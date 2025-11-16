package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateBuildCmd = &cobra.Command{
	Use:   "update-build",
	Short: "**This API works with the following fleet types:** EC2\n\nUpdates metadata in a build resource, including the build name and version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateBuildCmd).Standalone()

		gamelift_updateBuildCmd.Flags().String("build-id", "", "A unique identifier for the build to update.")
		gamelift_updateBuildCmd.Flags().String("name", "", "A descriptive label that is associated with a build.")
		gamelift_updateBuildCmd.Flags().String("version", "", "Version information that is associated with a build or script.")
		gamelift_updateBuildCmd.MarkFlagRequired("build-id")
	})
	gameliftCmd.AddCommand(gamelift_updateBuildCmd)
}
