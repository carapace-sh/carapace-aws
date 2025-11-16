package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_addTagsToVaultCmd = &cobra.Command{
	Use:   "add-tags-to-vault",
	Short: "This operation adds the specified tags to a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_addTagsToVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_addTagsToVaultCmd).Standalone()

		glacier_addTagsToVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_addTagsToVaultCmd.Flags().String("tags", "", "The tags to add to the vault.")
		glacier_addTagsToVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_addTagsToVaultCmd.MarkFlagRequired("account-id")
		glacier_addTagsToVaultCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_addTagsToVaultCmd)
}
