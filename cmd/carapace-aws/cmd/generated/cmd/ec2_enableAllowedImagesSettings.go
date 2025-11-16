package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableAllowedImagesSettingsCmd = &cobra.Command{
	Use:   "enable-allowed-images-settings",
	Short: "Enables Allowed AMIs for your account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableAllowedImagesSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableAllowedImagesSettingsCmd).Standalone()

		ec2_enableAllowedImagesSettingsCmd.Flags().String("allowed-images-settings-state", "", "Specify `enabled` to apply the image criteria specified by the Allowed AMIs settings.")
		ec2_enableAllowedImagesSettingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableAllowedImagesSettingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableAllowedImagesSettingsCmd.MarkFlagRequired("allowed-images-settings-state")
		ec2_enableAllowedImagesSettingsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableAllowedImagesSettingsCmd)
}
