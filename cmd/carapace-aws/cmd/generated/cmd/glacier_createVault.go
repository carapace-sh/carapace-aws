package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_createVaultCmd = &cobra.Command{
	Use:   "create-vault",
	Short: "This operation creates a new vault with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_createVaultCmd).Standalone()

	glacier_createVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
	glacier_createVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_createVaultCmd.MarkFlagRequired("account-id")
	glacier_createVaultCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_createVaultCmd)
}
