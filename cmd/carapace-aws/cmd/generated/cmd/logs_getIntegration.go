package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getIntegrationCmd = &cobra.Command{
	Use:   "get-integration",
	Short: "Returns information about one integration between CloudWatch Logs and OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getIntegrationCmd).Standalone()

	logs_getIntegrationCmd.Flags().String("integration-name", "", "The name of the integration that you want to find information about.")
	logs_getIntegrationCmd.MarkFlagRequired("integration-name")
	logsCmd.AddCommand(logs_getIntegrationCmd)
}
