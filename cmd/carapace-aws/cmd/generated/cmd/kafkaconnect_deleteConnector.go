package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_deleteConnectorCmd = &cobra.Command{
	Use:   "delete-connector",
	Short: "Deletes the specified connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_deleteConnectorCmd).Standalone()

	kafkaconnect_deleteConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector that you want to delete.")
	kafkaconnect_deleteConnectorCmd.Flags().String("current-version", "", "The current version of the connector that you want to delete.")
	kafkaconnect_deleteConnectorCmd.MarkFlagRequired("connector-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_deleteConnectorCmd)
}
