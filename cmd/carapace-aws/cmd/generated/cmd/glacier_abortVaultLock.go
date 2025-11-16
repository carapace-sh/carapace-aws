package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_abortVaultLockCmd = &cobra.Command{
	Use:   "abort-vault-lock",
	Short: "This operation aborts the vault locking process if the vault lock is not in the `Locked` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_abortVaultLockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_abortVaultLockCmd).Standalone()

		glacier_abortVaultLockCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
		glacier_abortVaultLockCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_abortVaultLockCmd.MarkFlagRequired("account-id")
		glacier_abortVaultLockCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_abortVaultLockCmd)
}
