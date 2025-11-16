package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Removes a user group from Amazon Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteGroupCmd).Standalone()

		quicksight_deleteGroupCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_deleteGroupCmd.Flags().String("group-name", "", "The name of the group that you want to delete.")
		quicksight_deleteGroupCmd.Flags().String("namespace", "", "The namespace of the group that you want to delete.")
		quicksight_deleteGroupCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteGroupCmd.MarkFlagRequired("group-name")
		quicksight_deleteGroupCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_deleteGroupCmd)
}
