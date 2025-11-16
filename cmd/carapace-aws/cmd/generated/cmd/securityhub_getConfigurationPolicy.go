package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getConfigurationPolicyCmd = &cobra.Command{
	Use:   "get-configuration-policy",
	Short: "Provides information about a configuration policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getConfigurationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getConfigurationPolicyCmd).Standalone()

		securityhub_getConfigurationPolicyCmd.Flags().String("identifier", "", "The Amazon Resource Name (ARN) or universally unique identifier (UUID) of the configuration policy.")
		securityhub_getConfigurationPolicyCmd.MarkFlagRequired("identifier")
	})
	securityhubCmd.AddCommand(securityhub_getConfigurationPolicyCmd)
}
