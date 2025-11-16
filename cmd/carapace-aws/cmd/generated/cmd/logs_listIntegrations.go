package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listIntegrationsCmd = &cobra.Command{
	Use:   "list-integrations",
	Short: "Returns a list of integrations between CloudWatch Logs and other services in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listIntegrationsCmd).Standalone()

	logs_listIntegrationsCmd.Flags().String("integration-name-prefix", "", "To limit the results to integrations that start with a certain name prefix, specify that name prefix here.")
	logs_listIntegrationsCmd.Flags().String("integration-status", "", "To limit the results to integrations with a certain status, specify that status here.")
	logs_listIntegrationsCmd.Flags().String("integration-type", "", "To limit the results to integrations of a certain type, specify that type here.")
	logsCmd.AddCommand(logs_listIntegrationsCmd)
}
