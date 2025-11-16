package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Deletes the integration between CloudWatch Logs and OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteIntegrationCmd).Standalone()

	logs_deleteIntegrationCmd.Flags().String("force", "", "Specify `true` to force the deletion of the integration even if vended logs dashboards currently exist.")
	logs_deleteIntegrationCmd.Flags().String("integration-name", "", "The name of the integration to delete.")
	logs_deleteIntegrationCmd.MarkFlagRequired("integration-name")
	logsCmd.AddCommand(logs_deleteIntegrationCmd)
}
