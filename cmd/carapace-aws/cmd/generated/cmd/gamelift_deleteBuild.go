package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteBuildCmd = &cobra.Command{
	Use:   "delete-build",
	Short: "**This API works with the following fleet types:** EC2\n\nDeletes a build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteBuildCmd).Standalone()

	gamelift_deleteBuildCmd.Flags().String("build-id", "", "A unique identifier for the build to delete.")
	gamelift_deleteBuildCmd.MarkFlagRequired("build-id")
	gameliftCmd.AddCommand(gamelift_deleteBuildCmd)
}
