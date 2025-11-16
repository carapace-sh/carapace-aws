package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group so that all users and sub groups that belong to the group can no longer access documents only available to that group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteGroupCmd).Standalone()

	qbusiness_deleteGroupCmd.Flags().String("application-id", "", "The identifier of the application in which the group mapping belongs.")
	qbusiness_deleteGroupCmd.Flags().String("data-source-id", "", "The identifier of the data source linked to the group")
	qbusiness_deleteGroupCmd.Flags().String("group-name", "", "The name of the group you want to delete.")
	qbusiness_deleteGroupCmd.Flags().String("index-id", "", "The identifier of the index you want to delete the group from.")
	qbusiness_deleteGroupCmd.MarkFlagRequired("application-id")
	qbusiness_deleteGroupCmd.MarkFlagRequired("group-name")
	qbusiness_deleteGroupCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_deleteGroupCmd)
}
