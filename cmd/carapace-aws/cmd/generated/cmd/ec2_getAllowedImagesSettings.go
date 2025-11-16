package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getAllowedImagesSettingsCmd = &cobra.Command{
	Use:   "get-allowed-images-settings",
	Short: "Gets the current state of the Allowed AMIs setting and the list of Allowed AMIs criteria at the account level in the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getAllowedImagesSettingsCmd).Standalone()

	ec2_getAllowedImagesSettingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getAllowedImagesSettingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getAllowedImagesSettingsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getAllowedImagesSettingsCmd)
}
