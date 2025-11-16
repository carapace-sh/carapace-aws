package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAccountSettingsCmd = &cobra.Command{
	Use:   "describe-account-settings",
	Short: "Describes the settings that were used when your Quick Sight subscription was first created in this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAccountSettingsCmd).Standalone()

	quicksight_describeAccountSettingsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the settings that you want to list.")
	quicksight_describeAccountSettingsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeAccountSettingsCmd)
}
