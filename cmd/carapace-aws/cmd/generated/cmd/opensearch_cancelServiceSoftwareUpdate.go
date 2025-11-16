package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_cancelServiceSoftwareUpdateCmd = &cobra.Command{
	Use:   "cancel-service-software-update",
	Short: "Cancels a scheduled service software update for an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_cancelServiceSoftwareUpdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_cancelServiceSoftwareUpdateCmd).Standalone()

		opensearch_cancelServiceSoftwareUpdateCmd.Flags().String("domain-name", "", "Name of the OpenSearch Service domain that you want to cancel the service software update on.")
		opensearch_cancelServiceSoftwareUpdateCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_cancelServiceSoftwareUpdateCmd)
}
