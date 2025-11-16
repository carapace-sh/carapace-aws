package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_setTokenVaultCmkCmd = &cobra.Command{
	Use:   "set-token-vault-cmk",
	Short: "Sets the customer master key (CMK) for a token vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_setTokenVaultCmkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_setTokenVaultCmkCmd).Standalone()

		bedrockAgentcoreControl_setTokenVaultCmkCmd.Flags().String("kms-configuration", "", "The KMS configuration for the token vault, including the key type and KMS key ARN.")
		bedrockAgentcoreControl_setTokenVaultCmkCmd.Flags().String("token-vault-id", "", "The unique identifier of the token vault to update.")
		bedrockAgentcoreControl_setTokenVaultCmkCmd.MarkFlagRequired("kms-configuration")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_setTokenVaultCmkCmd)
}
