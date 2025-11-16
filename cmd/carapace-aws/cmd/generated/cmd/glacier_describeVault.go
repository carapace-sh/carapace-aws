package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_describeVaultCmd = &cobra.Command{
	Use:   "describe-vault",
	Short: "This operation returns information about a vault, including the vault's Amazon Resource Name (ARN), the date the vault was created, the number of archives it contains, and the total size of all the archives in the vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_describeVaultCmd).Standalone()

	glacier_describeVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_describeVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_describeVaultCmd.MarkFlagRequired("account-id")
	glacier_describeVaultCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_describeVaultCmd)
}
