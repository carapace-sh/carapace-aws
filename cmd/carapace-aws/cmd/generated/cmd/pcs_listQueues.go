package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_listQueuesCmd = &cobra.Command{
	Use:   "list-queues",
	Short: "Returns a list of all queues associated with a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_listQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_listQueuesCmd).Standalone()

		pcs_listQueuesCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to list queues for.")
		pcs_listQueuesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		pcs_listQueuesCmd.Flags().String("next-token", "", "The value of `nextToken` is a unique pagination token for each page of results returned.")
		pcs_listQueuesCmd.MarkFlagRequired("cluster-identifier")
	})
	pcsCmd.AddCommand(pcs_listQueuesCmd)
}
