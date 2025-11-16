package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_setVaultAccessPolicyCmd = &cobra.Command{
	Use:   "set-vault-access-policy",
	Short: "This operation configures an access policy for a vault and will overwrite an existing policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_setVaultAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_setVaultAccessPolicyCmd).Standalone()

		glacier_setVaultAccessPolicyCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_setVaultAccessPolicyCmd.Flags().String("policy", "", "The vault access policy as a JSON string.")
		glacier_setVaultAccessPolicyCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_setVaultAccessPolicyCmd.MarkFlagRequired("account-id")
		glacier_setVaultAccessPolicyCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_setVaultAccessPolicyCmd)
}
