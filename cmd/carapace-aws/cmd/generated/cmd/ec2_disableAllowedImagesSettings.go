package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableAllowedImagesSettingsCmd = &cobra.Command{
	Use:   "disable-allowed-images-settings",
	Short: "Disables Allowed AMIs for your account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableAllowedImagesSettingsCmd).Standalone()

	ec2_disableAllowedImagesSettingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableAllowedImagesSettingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableAllowedImagesSettingsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableAllowedImagesSettingsCmd)
}
