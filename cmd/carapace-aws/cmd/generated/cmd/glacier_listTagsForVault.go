package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listTagsForVaultCmd = &cobra.Command{
	Use:   "list-tags-for-vault",
	Short: "This operation lists all the tags attached to a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listTagsForVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_listTagsForVaultCmd).Standalone()

		glacier_listTagsForVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_listTagsForVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_listTagsForVaultCmd.MarkFlagRequired("account-id")
		glacier_listTagsForVaultCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_listTagsForVaultCmd)
}
