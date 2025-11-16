package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateConfigurationPolicyCmd = &cobra.Command{
	Use:   "update-configuration-policy",
	Short: "Updates a configuration policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateConfigurationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateConfigurationPolicyCmd).Standalone()

		securityhub_updateConfigurationPolicyCmd.Flags().String("configuration-policy", "", "An object that defines how Security Hub is configured.")
		securityhub_updateConfigurationPolicyCmd.Flags().String("description", "", "The description of the configuration policy.")
		securityhub_updateConfigurationPolicyCmd.Flags().String("identifier", "", "The Amazon Resource Name (ARN) or universally unique identifier (UUID) of the configuration policy.")
		securityhub_updateConfigurationPolicyCmd.Flags().String("name", "", "The name of the configuration policy.")
		securityhub_updateConfigurationPolicyCmd.Flags().String("updated-reason", "", "The reason for updating the configuration policy.")
		securityhub_updateConfigurationPolicyCmd.MarkFlagRequired("identifier")
	})
	securityhubCmd.AddCommand(securityhub_updateConfigurationPolicyCmd)
}
