package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putRemediationConfigurationsCmd = &cobra.Command{
	Use:   "put-remediation-configurations",
	Short: "Adds or updates the remediation configuration with a specific Config rule with the selected target or action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putRemediationConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putRemediationConfigurationsCmd).Standalone()

		config_putRemediationConfigurationsCmd.Flags().String("remediation-configurations", "", "A list of remediation configuration objects.")
		config_putRemediationConfigurationsCmd.MarkFlagRequired("remediation-configurations")
	})
	configCmd.AddCommand(config_putRemediationConfigurationsCmd)
}
