package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_startServiceSoftwareUpdateCmd = &cobra.Command{
	Use:   "start-service-software-update",
	Short: "Schedules a service software update for an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_startServiceSoftwareUpdateCmd).Standalone()

	opensearch_startServiceSoftwareUpdateCmd.Flags().String("desired-start-time", "", "The Epoch timestamp when you want the service software update to start.")
	opensearch_startServiceSoftwareUpdateCmd.Flags().String("domain-name", "", "The name of the domain that you want to update to the latest service software.")
	opensearch_startServiceSoftwareUpdateCmd.Flags().String("schedule-at", "", "When to start the service software update.")
	opensearch_startServiceSoftwareUpdateCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_startServiceSoftwareUpdateCmd)
}
