package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_listConnectorOperationsCmd = &cobra.Command{
	Use:   "list-connector-operations",
	Short: "Lists information about a connector's operation(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_listConnectorOperationsCmd).Standalone()

	kafkaconnect_listConnectorOperationsCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector for which to list operations.")
	kafkaconnect_listConnectorOperationsCmd.Flags().String("max-results", "", "Maximum number of connector operations to fetch in one get request.")
	kafkaconnect_listConnectorOperationsCmd.Flags().String("next-token", "", "If the response is truncated, it includes a NextToken.")
	kafkaconnect_listConnectorOperationsCmd.MarkFlagRequired("connector-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_listConnectorOperationsCmd)
}
