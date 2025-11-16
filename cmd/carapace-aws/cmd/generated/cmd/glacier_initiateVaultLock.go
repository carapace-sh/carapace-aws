package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_initiateVaultLockCmd = &cobra.Command{
	Use:   "initiate-vault-lock",
	Short: "This operation initiates the vault locking process by doing the following:\n\n- Installing a vault lock policy on the specified vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_initiateVaultLockCmd).Standalone()

	glacier_initiateVaultLockCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
	glacier_initiateVaultLockCmd.Flags().String("policy", "", "The vault lock policy as a JSON string, which uses \"\\\\\" as an escape character.")
	glacier_initiateVaultLockCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_initiateVaultLockCmd.MarkFlagRequired("account-id")
	glacier_initiateVaultLockCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_initiateVaultLockCmd)
}
