package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeLoggingConfigurationCmd = &cobra.Command{
	Use:   "describe-logging-configuration",
	Short: "Returns complete information about the current rules and alerting logging configuration of the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeLoggingConfigurationCmd).Standalone()

		amp_describeLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace to describe the logging configuration for.")
		amp_describeLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeLoggingConfigurationCmd)
}
