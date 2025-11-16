package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeQueryLoggingConfigurationCmd = &cobra.Command{
	Use:   "describe-query-logging-configuration",
	Short: "Retrieves the details of the query logging configuration for the specified workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeQueryLoggingConfigurationCmd).Standalone()

	amp_describeQueryLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace for which to retrieve the query logging configuration.")
	amp_describeQueryLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_describeQueryLoggingConfigurationCmd)
}
