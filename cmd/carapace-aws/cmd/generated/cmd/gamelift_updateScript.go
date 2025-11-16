package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateScriptCmd = &cobra.Command{
	Use:   "update-script",
	Short: "**This API works with the following fleet types:** EC2\n\nUpdates Realtime script metadata and content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateScriptCmd).Standalone()

	gamelift_updateScriptCmd.Flags().String("name", "", "A descriptive label that is associated with a script.")
	gamelift_updateScriptCmd.Flags().String("script-id", "", "A unique identifier for the Realtime script to update.")
	gamelift_updateScriptCmd.Flags().String("storage-location", "", "The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored.")
	gamelift_updateScriptCmd.Flags().String("version", "", "Version information that is associated with a build or script.")
	gamelift_updateScriptCmd.Flags().String("zip-file", "", "A data object containing your Realtime scripts and dependencies as a zip file.")
	gamelift_updateScriptCmd.MarkFlagRequired("script-id")
	gameliftCmd.AddCommand(gamelift_updateScriptCmd)
}
