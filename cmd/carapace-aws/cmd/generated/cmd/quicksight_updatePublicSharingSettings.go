package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updatePublicSharingSettingsCmd = &cobra.Command{
	Use:   "update-public-sharing-settings",
	Short: "This API controls public sharing settings for your entire Quick Sight account, affecting data security and access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updatePublicSharingSettingsCmd).Standalone()

	quicksight_updatePublicSharingSettingsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID associated with your Amazon Quick Sight subscription.")
	quicksight_updatePublicSharingSettingsCmd.Flags().Bool("no-public-sharing-enabled", false, "A Boolean value that indicates whether public sharing is turned on for an Quick Suite account.")
	quicksight_updatePublicSharingSettingsCmd.Flags().Bool("public-sharing-enabled", false, "A Boolean value that indicates whether public sharing is turned on for an Quick Suite account.")
	quicksight_updatePublicSharingSettingsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updatePublicSharingSettingsCmd.Flag("no-public-sharing-enabled").Hidden = true
	quicksightCmd.AddCommand(quicksight_updatePublicSharingSettingsCmd)
}
