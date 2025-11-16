package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_updateAccountConfigurationCmd = &cobra.Command{
	Use:   "update-account-configuration",
	Short: "Use to update the encryption configuration for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_updateAccountConfigurationCmd).Standalone()

	codeguruSecurity_updateAccountConfigurationCmd.Flags().String("encryption-config", "", "The customer-managed KMS key ARN you want to use for encryption.")
	codeguruSecurity_updateAccountConfigurationCmd.MarkFlagRequired("encryption-config")
	codeguruSecurityCmd.AddCommand(codeguruSecurity_updateAccountConfigurationCmd)
}
