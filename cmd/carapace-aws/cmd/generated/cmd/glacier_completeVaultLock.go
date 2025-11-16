package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_completeVaultLockCmd = &cobra.Command{
	Use:   "complete-vault-lock",
	Short: "This operation completes the vault locking process by transitioning the vault lock from the `InProgress` state to the `Locked` state, which causes the vault lock policy to become unchangeable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_completeVaultLockCmd).Standalone()

	glacier_completeVaultLockCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
	glacier_completeVaultLockCmd.Flags().String("lock-id", "", "The `lockId` value is the lock ID obtained from a [InitiateVaultLock]() request.")
	glacier_completeVaultLockCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_completeVaultLockCmd.MarkFlagRequired("account-id")
	glacier_completeVaultLockCmd.MarkFlagRequired("lock-id")
	glacier_completeVaultLockCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_completeVaultLockCmd)
}
