package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Returns a list of all the connectors in this account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_listConnectorsCmd).Standalone()

	kafkaconnect_listConnectorsCmd.Flags().String("connector-name-prefix", "", "The name prefix that you want to use to search for and list connectors.")
	kafkaconnect_listConnectorsCmd.Flags().String("max-results", "", "The maximum number of connectors to list in one response.")
	kafkaconnect_listConnectorsCmd.Flags().String("next-token", "", "If the response of a ListConnectors operation is truncated, it will include a NextToken.")
	kafkaconnectCmd.AddCommand(kafkaconnect_listConnectorsCmd)
}
