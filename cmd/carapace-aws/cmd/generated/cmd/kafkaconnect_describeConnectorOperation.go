package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_describeConnectorOperationCmd = &cobra.Command{
	Use:   "describe-connector-operation",
	Short: "Returns information about the specified connector's operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_describeConnectorOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_describeConnectorOperationCmd).Standalone()

		kafkaconnect_describeConnectorOperationCmd.Flags().String("connector-operation-arn", "", "ARN of the connector operation to be described.")
		kafkaconnect_describeConnectorOperationCmd.MarkFlagRequired("connector-operation-arn")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_describeConnectorOperationCmd)
}
