package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeRemediationConfigurationsCmd = &cobra.Command{
	Use:   "describe-remediation-configurations",
	Short: "Returns the details of one or more remediation configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeRemediationConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeRemediationConfigurationsCmd).Standalone()

		config_describeRemediationConfigurationsCmd.Flags().String("config-rule-names", "", "A list of Config rule names of remediation configurations for which you want details.")
		config_describeRemediationConfigurationsCmd.MarkFlagRequired("config-rule-names")
	})
	configCmd.AddCommand(config_describeRemediationConfigurationsCmd)
}
