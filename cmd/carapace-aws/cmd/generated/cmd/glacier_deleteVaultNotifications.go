package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_deleteVaultNotificationsCmd = &cobra.Command{
	Use:   "delete-vault-notifications",
	Short: "This operation deletes the notification configuration set for a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_deleteVaultNotificationsCmd).Standalone()

	glacier_deleteVaultNotificationsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_deleteVaultNotificationsCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_deleteVaultNotificationsCmd.MarkFlagRequired("account-id")
	glacier_deleteVaultNotificationsCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_deleteVaultNotificationsCmd)
}
