package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_setSmbguestPasswordCmd = &cobra.Command{
	Use:   "set-smbguest-password",
	Short: "Sets the password for the guest user `smbguest`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_setSmbguestPasswordCmd).Standalone()

	storagegateway_setSmbguestPasswordCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the S3 File Gateway the SMB file share is associated with.")
	storagegateway_setSmbguestPasswordCmd.Flags().String("password", "", "The password that you want to set for your SMB server.")
	storagegateway_setSmbguestPasswordCmd.MarkFlagRequired("gateway-arn")
	storagegateway_setSmbguestPasswordCmd.MarkFlagRequired("password")
	storagegatewayCmd.AddCommand(storagegateway_setSmbguestPasswordCmd)
}
