package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteQueryLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-query-logging-configuration",
	Short: "Deletes the query logging configuration for the specified workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteQueryLoggingConfigurationCmd).Standalone()

	amp_deleteQueryLoggingConfigurationCmd.Flags().String("client-token", "", "(Optional) A unique, case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	amp_deleteQueryLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace from which to delete the query logging configuration.")
	amp_deleteQueryLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_deleteQueryLoggingConfigurationCmd)
}
