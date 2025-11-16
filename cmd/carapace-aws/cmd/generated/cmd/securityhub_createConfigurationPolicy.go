package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createConfigurationPolicyCmd = &cobra.Command{
	Use:   "create-configuration-policy",
	Short: "Creates a configuration policy with the defined configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createConfigurationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createConfigurationPolicyCmd).Standalone()

		securityhub_createConfigurationPolicyCmd.Flags().String("configuration-policy", "", "An object that defines how Security Hub is configured.")
		securityhub_createConfigurationPolicyCmd.Flags().String("description", "", "The description of the configuration policy.")
		securityhub_createConfigurationPolicyCmd.Flags().String("name", "", "The name of the configuration policy.")
		securityhub_createConfigurationPolicyCmd.Flags().String("tags", "", "User-defined tags associated with a configuration policy.")
		securityhub_createConfigurationPolicyCmd.MarkFlagRequired("configuration-policy")
		securityhub_createConfigurationPolicyCmd.MarkFlagRequired("name")
	})
	securityhubCmd.AddCommand(securityhub_createConfigurationPolicyCmd)
}
