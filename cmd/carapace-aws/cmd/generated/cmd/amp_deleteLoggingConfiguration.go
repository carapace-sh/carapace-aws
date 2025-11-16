package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-logging-configuration",
	Short: "Deletes the rules and alerting logging configuration for a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteLoggingConfigurationCmd).Standalone()

	amp_deleteLoggingConfigurationCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
	amp_deleteLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace containing the logging configuration to delete.")
	amp_deleteLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_deleteLoggingConfigurationCmd)
}
