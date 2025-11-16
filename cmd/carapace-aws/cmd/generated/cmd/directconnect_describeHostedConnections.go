package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeHostedConnectionsCmd = &cobra.Command{
	Use:   "describe-hosted-connections",
	Short: "Lists the hosted connections that have been provisioned on the specified interconnect or link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeHostedConnectionsCmd).Standalone()

	directconnect_describeHostedConnectionsCmd.Flags().String("connection-id", "", "The ID of the interconnect or LAG.")
	directconnect_describeHostedConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeHostedConnectionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	directconnect_describeHostedConnectionsCmd.MarkFlagRequired("connection-id")
	directconnectCmd.AddCommand(directconnect_describeHostedConnectionsCmd)
}
