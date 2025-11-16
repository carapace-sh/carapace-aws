package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateQueryLoggingConfigurationCmd = &cobra.Command{
	Use:   "update-query-logging-configuration",
	Short: "Updates the query logging configuration for the specified workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateQueryLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_updateQueryLoggingConfigurationCmd).Standalone()

		amp_updateQueryLoggingConfigurationCmd.Flags().String("client-token", "", "(Optional) A unique, case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		amp_updateQueryLoggingConfigurationCmd.Flags().String("destinations", "", "The destinations where query logs will be sent.")
		amp_updateQueryLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace for which to update the query logging configuration.")
		amp_updateQueryLoggingConfigurationCmd.MarkFlagRequired("destinations")
		amp_updateQueryLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_updateQueryLoggingConfigurationCmd)
}
