package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createScriptCmd = &cobra.Command{
	Use:   "create-script",
	Short: "**This API works with the following fleet types:** EC2, Anywhere\n\nCreates a new script record for your Amazon GameLift Servers Realtime script.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createScriptCmd).Standalone()

	gamelift_createScriptCmd.Flags().String("name", "", "A descriptive label that is associated with a script.")
	gamelift_createScriptCmd.Flags().String("storage-location", "", "The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored.")
	gamelift_createScriptCmd.Flags().String("tags", "", "A list of labels to assign to the new script resource.")
	gamelift_createScriptCmd.Flags().String("version", "", "Version information that is associated with a build or script.")
	gamelift_createScriptCmd.Flags().String("zip-file", "", "A data object containing your Realtime scripts and dependencies as a zip file.")
	gameliftCmd.AddCommand(gamelift_createScriptCmd)
}
