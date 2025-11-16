package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Updates the Amazon Quick Sight settings in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateAccountSettingsCmd).Standalone()

		quicksight_updateAccountSettingsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the Quick Sight settings that you want to list.")
		quicksight_updateAccountSettingsCmd.Flags().String("default-namespace", "", "The default namespace for this Amazon Web Services account.")
		quicksight_updateAccountSettingsCmd.Flags().Bool("no-termination-protection-enabled", false, "A boolean value that determines whether or not an Quick Sight account can be deleted.")
		quicksight_updateAccountSettingsCmd.Flags().String("notification-email", "", "The email address that you want Quick Sight to send notifications to regarding your Amazon Web Services account or Quick Sight subscription.")
		quicksight_updateAccountSettingsCmd.Flags().Bool("termination-protection-enabled", false, "A boolean value that determines whether or not an Quick Sight account can be deleted.")
		quicksight_updateAccountSettingsCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateAccountSettingsCmd.MarkFlagRequired("default-namespace")
		quicksight_updateAccountSettingsCmd.Flag("no-termination-protection-enabled").Hidden = true
	})
	quicksightCmd.AddCommand(quicksight_updateAccountSettingsCmd)
}
