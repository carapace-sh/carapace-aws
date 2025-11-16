package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceImageCriteriaInAllowedImagesSettingsCmd = &cobra.Command{
	Use:   "replace-image-criteria-in-allowed-images-settings",
	Short: "Sets or replaces the criteria for Allowed AMIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceImageCriteriaInAllowedImagesSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_replaceImageCriteriaInAllowedImagesSettingsCmd).Standalone()

		ec2_replaceImageCriteriaInAllowedImagesSettingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceImageCriteriaInAllowedImagesSettingsCmd.Flags().String("image-criteria", "", "The list of criteria that are evaluated to determine whether AMIs are discoverable and usable in the account in the specified Amazon Web Services Region.")
		ec2_replaceImageCriteriaInAllowedImagesSettingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceImageCriteriaInAllowedImagesSettingsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_replaceImageCriteriaInAllowedImagesSettingsCmd)
}
