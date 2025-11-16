package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createQueryLoggingConfigurationCmd = &cobra.Command{
	Use:   "create-query-logging-configuration",
	Short: "Creates a query logging configuration for the specified workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createQueryLoggingConfigurationCmd).Standalone()

	amp_createQueryLoggingConfigurationCmd.Flags().String("client-token", "", "(Optional) A unique, case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	amp_createQueryLoggingConfigurationCmd.Flags().String("destinations", "", "The destinations where query logs will be sent.")
	amp_createQueryLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace for which to create the query logging configuration.")
	amp_createQueryLoggingConfigurationCmd.MarkFlagRequired("destinations")
	amp_createQueryLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_createQueryLoggingConfigurationCmd)
}
