package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_getVaultAccessPolicyCmd = &cobra.Command{
	Use:   "get-vault-access-policy",
	Short: "This operation retrieves the `access-policy` subresource set on the vault; for more information on setting this subresource, see [Set Vault Access Policy (PUT access-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-SetVaultAccessPolicy.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_getVaultAccessPolicyCmd).Standalone()

	glacier_getVaultAccessPolicyCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_getVaultAccessPolicyCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_getVaultAccessPolicyCmd.MarkFlagRequired("account-id")
	glacier_getVaultAccessPolicyCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_getVaultAccessPolicyCmd)
}
