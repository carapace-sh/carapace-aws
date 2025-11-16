package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_updateConnectorCmd = &cobra.Command{
	Use:   "update-connector",
	Short: "Updates the specified connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_updateConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_updateConnectorCmd).Standalone()

		kafkaconnect_updateConnectorCmd.Flags().String("capacity", "", "The target capacity.")
		kafkaconnect_updateConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector that you want to update.")
		kafkaconnect_updateConnectorCmd.Flags().String("connector-configuration", "", "A map of keys to values that represent the configuration for the connector.")
		kafkaconnect_updateConnectorCmd.Flags().String("current-version", "", "The current version of the connector that you want to update.")
		kafkaconnect_updateConnectorCmd.MarkFlagRequired("connector-arn")
		kafkaconnect_updateConnectorCmd.MarkFlagRequired("current-version")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_updateConnectorCmd)
}
