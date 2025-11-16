package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteDomainConfigurationCmd = &cobra.Command{
	Use:   "delete-domain-configuration",
	Short: "Deletes the specified domain configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteDomainConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteDomainConfigurationCmd).Standalone()

		iot_deleteDomainConfigurationCmd.Flags().String("domain-configuration-name", "", "The name of the domain configuration to be deleted.")
		iot_deleteDomainConfigurationCmd.MarkFlagRequired("domain-configuration-name")
	})
	iotCmd.AddCommand(iot_deleteDomainConfigurationCmd)
}
