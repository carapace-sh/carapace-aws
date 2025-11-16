package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Describes a group by group name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getGroupCmd).Standalone()

		qbusiness_getGroupCmd.Flags().String("application-id", "", "The identifier of the application id the group is attached to.")
		qbusiness_getGroupCmd.Flags().String("data-source-id", "", "The identifier of the data source the group is attached to.")
		qbusiness_getGroupCmd.Flags().String("group-name", "", "The name of the group.")
		qbusiness_getGroupCmd.Flags().String("index-id", "", "The identifier of the index the group is attached to.")
		qbusiness_getGroupCmd.MarkFlagRequired("application-id")
		qbusiness_getGroupCmd.MarkFlagRequired("group-name")
		qbusiness_getGroupCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getGroupCmd)
}
