package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Changes a group description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateGroupCmd).Standalone()

	quicksight_updateGroupCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_updateGroupCmd.Flags().String("description", "", "The description for the group that you want to update.")
	quicksight_updateGroupCmd.Flags().String("group-name", "", "The name of the group that you want to update.")
	quicksight_updateGroupCmd.Flags().String("namespace", "", "The namespace of the group that you want to update.")
	quicksight_updateGroupCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateGroupCmd.MarkFlagRequired("group-name")
	quicksight_updateGroupCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_updateGroupCmd)
}
