package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteRemediationConfigurationCmd = &cobra.Command{
	Use:   "delete-remediation-configuration",
	Short: "Deletes the remediation configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteRemediationConfigurationCmd).Standalone()

	config_deleteRemediationConfigurationCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want to delete remediation configuration.")
	config_deleteRemediationConfigurationCmd.Flags().String("resource-type", "", "The type of a resource.")
	config_deleteRemediationConfigurationCmd.MarkFlagRequired("config-rule-name")
	configCmd.AddCommand(config_deleteRemediationConfigurationCmd)
}
