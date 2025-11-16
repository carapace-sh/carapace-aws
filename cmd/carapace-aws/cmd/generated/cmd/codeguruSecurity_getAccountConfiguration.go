package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_getAccountConfigurationCmd = &cobra.Command{
	Use:   "get-account-configuration",
	Short: "Use to get the encryption configuration for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_getAccountConfigurationCmd).Standalone()

	codeguruSecurityCmd.AddCommand(codeguruSecurity_getAccountConfigurationCmd)
}
