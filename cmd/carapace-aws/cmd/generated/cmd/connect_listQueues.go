package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listQueuesCmd = &cobra.Command{
	Use:   "list-queues",
	Short: "Provides information about the queues for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listQueuesCmd).Standalone()

		connect_listQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listQueuesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listQueuesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listQueuesCmd.Flags().String("queue-types", "", "The type of queue.")
		connect_listQueuesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listQueuesCmd)
}
