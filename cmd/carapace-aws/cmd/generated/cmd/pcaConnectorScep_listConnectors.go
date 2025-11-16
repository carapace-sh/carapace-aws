package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Lists the connectors belonging to your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_listConnectorsCmd).Standalone()

	pcaConnectorScep_listConnectorsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Connector for SCEP to return for this request.")
	pcaConnectorScep_listConnectorsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Connector for SCEP returns a `NextToken` value in the response.")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_listConnectorsCmd)
}
