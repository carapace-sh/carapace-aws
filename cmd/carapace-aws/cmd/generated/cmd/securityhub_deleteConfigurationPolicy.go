package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteConfigurationPolicyCmd = &cobra.Command{
	Use:   "delete-configuration-policy",
	Short: "Deletes a configuration policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteConfigurationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_deleteConfigurationPolicyCmd).Standalone()

		securityhub_deleteConfigurationPolicyCmd.Flags().String("identifier", "", "The Amazon Resource Name (ARN) or universally unique identifier (UUID) of the configuration policy.")
		securityhub_deleteConfigurationPolicyCmd.MarkFlagRequired("identifier")
	})
	securityhubCmd.AddCommand(securityhub_deleteConfigurationPolicyCmd)
}
