package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_getVaultNotificationsCmd = &cobra.Command{
	Use:   "get-vault-notifications",
	Short: "This operation retrieves the `notification-configuration` subresource of the specified vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_getVaultNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_getVaultNotificationsCmd).Standalone()

		glacier_getVaultNotificationsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_getVaultNotificationsCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_getVaultNotificationsCmd.MarkFlagRequired("account-id")
		glacier_getVaultNotificationsCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_getVaultNotificationsCmd)
}
