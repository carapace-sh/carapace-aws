package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putIntegrationCmd = &cobra.Command{
	Use:   "put-integration",
	Short: "Creates an integration between CloudWatch Logs and another service in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putIntegrationCmd).Standalone()

		logs_putIntegrationCmd.Flags().String("integration-name", "", "A name for the integration.")
		logs_putIntegrationCmd.Flags().String("integration-type", "", "The type of integration.")
		logs_putIntegrationCmd.Flags().String("resource-config", "", "A structure that contains configuration information for the integration that you are creating.")
		logs_putIntegrationCmd.MarkFlagRequired("integration-name")
		logs_putIntegrationCmd.MarkFlagRequired("integration-type")
		logs_putIntegrationCmd.MarkFlagRequired("resource-config")
	})
	logsCmd.AddCommand(logs_putIntegrationCmd)
}
