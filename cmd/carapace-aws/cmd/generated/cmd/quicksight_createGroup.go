package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Use the `CreateGroup` operation to create a group in Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createGroupCmd).Standalone()

	quicksight_createGroupCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_createGroupCmd.Flags().String("description", "", "A description for the group that you want to create.")
	quicksight_createGroupCmd.Flags().String("group-name", "", "A name for the group that you want to create.")
	quicksight_createGroupCmd.Flags().String("namespace", "", "The namespace that you want the group to be a part of.")
	quicksight_createGroupCmd.MarkFlagRequired("aws-account-id")
	quicksight_createGroupCmd.MarkFlagRequired("group-name")
	quicksight_createGroupCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_createGroupCmd)
}
