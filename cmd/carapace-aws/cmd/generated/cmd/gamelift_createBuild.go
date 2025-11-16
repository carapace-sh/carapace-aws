package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createBuildCmd = &cobra.Command{
	Use:   "create-build",
	Short: "**This API works with the following fleet types:** EC2, Anywhere\n\nCreates a new Amazon GameLift Servers build resource for your game server binary files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createBuildCmd).Standalone()

		gamelift_createBuildCmd.Flags().String("name", "", "A descriptive label that is associated with a build.")
		gamelift_createBuildCmd.Flags().String("operating-system", "", "The operating system that your game server binaries run on.")
		gamelift_createBuildCmd.Flags().String("server-sdk-version", "", "A server SDK version you used when integrating your game server build with Amazon GameLift Servers.")
		gamelift_createBuildCmd.Flags().String("storage-location", "", "Information indicating where your game build files are stored.")
		gamelift_createBuildCmd.Flags().String("tags", "", "A list of labels to assign to the new build resource.")
		gamelift_createBuildCmd.Flags().String("version", "", "Version information that is associated with a build or script.")
	})
	gameliftCmd.AddCommand(gamelift_createBuildCmd)
}
