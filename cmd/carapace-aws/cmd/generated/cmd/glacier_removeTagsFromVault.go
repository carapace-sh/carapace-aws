package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_removeTagsFromVaultCmd = &cobra.Command{
	Use:   "remove-tags-from-vault",
	Short: "This operation removes one or more tags from the set of tags attached to a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_removeTagsFromVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_removeTagsFromVaultCmd).Standalone()

		glacier_removeTagsFromVaultCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_removeTagsFromVaultCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		glacier_removeTagsFromVaultCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_removeTagsFromVaultCmd.MarkFlagRequired("account-id")
		glacier_removeTagsFromVaultCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_removeTagsFromVaultCmd)
}
