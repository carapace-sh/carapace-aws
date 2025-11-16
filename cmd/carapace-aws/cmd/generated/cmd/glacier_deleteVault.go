package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_deleteVaultCmd = &cobra.Command{
	Use:   "delete-vault",
	Short: "This operation deletes a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_deleteVaultCmd).Standalone()

	glacier_deleteVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_deleteVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_deleteVaultCmd.MarkFlagRequired("account-id")
	glacier_deleteVaultCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_deleteVaultCmd)
}
