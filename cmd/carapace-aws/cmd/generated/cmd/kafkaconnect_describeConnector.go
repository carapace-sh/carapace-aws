package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_describeConnectorCmd = &cobra.Command{
	Use:   "describe-connector",
	Short: "Returns summary information about the connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_describeConnectorCmd).Standalone()

	kafkaconnect_describeConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector that you want to describe.")
	kafkaconnect_describeConnectorCmd.MarkFlagRequired("connector-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_describeConnectorCmd)
}
