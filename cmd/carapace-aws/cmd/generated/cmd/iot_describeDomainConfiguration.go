package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeDomainConfigurationCmd = &cobra.Command{
	Use:   "describe-domain-configuration",
	Short: "Gets summary information about a domain configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeDomainConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeDomainConfigurationCmd).Standalone()

		iot_describeDomainConfigurationCmd.Flags().String("domain-configuration-name", "", "The name of the domain configuration.")
		iot_describeDomainConfigurationCmd.MarkFlagRequired("domain-configuration-name")
	})
	iotCmd.AddCommand(iot_describeDomainConfigurationCmd)
}
