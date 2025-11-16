package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_getVaultLockCmd = &cobra.Command{
	Use:   "get-vault-lock",
	Short: "This operation retrieves the following attributes from the `lock-policy` subresource set on the specified vault:\n\n- The vault lock policy set on the vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_getVaultLockCmd).Standalone()

	glacier_getVaultLockCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_getVaultLockCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_getVaultLockCmd.MarkFlagRequired("account-id")
	glacier_getVaultLockCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_getVaultLockCmd)
}
