package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_setVaultNotificationsCmd = &cobra.Command{
	Use:   "set-vault-notifications",
	Short: "This operation configures notifications that will be sent when specific events happen to a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_setVaultNotificationsCmd).Standalone()

	glacier_setVaultNotificationsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_setVaultNotificationsCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_setVaultNotificationsCmd.Flags().String("vault-notification-config", "", "Provides options for specifying notification configuration.")
	glacier_setVaultNotificationsCmd.MarkFlagRequired("account-id")
	glacier_setVaultNotificationsCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_setVaultNotificationsCmd)
}
