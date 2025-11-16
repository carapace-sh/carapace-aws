package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_deleteVaultAccessPolicyCmd = &cobra.Command{
	Use:   "delete-vault-access-policy",
	Short: "This operation deletes the access policy associated with the specified vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_deleteVaultAccessPolicyCmd).Standalone()

	glacier_deleteVaultAccessPolicyCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_deleteVaultAccessPolicyCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_deleteVaultAccessPolicyCmd.MarkFlagRequired("account-id")
	glacier_deleteVaultAccessPolicyCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_deleteVaultAccessPolicyCmd)
}
