package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_putLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-logging-configuration",
	Short: "Enables the specified [LoggingConfiguration](), to start logging from a web ACL, according to the configuration provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_putLoggingConfigurationCmd).Standalone()

	wafv2_putLoggingConfigurationCmd.Flags().String("logging-configuration", "", "")
	wafv2_putLoggingConfigurationCmd.MarkFlagRequired("logging-configuration")
	wafv2Cmd.AddCommand(wafv2_putLoggingConfigurationCmd)
}
