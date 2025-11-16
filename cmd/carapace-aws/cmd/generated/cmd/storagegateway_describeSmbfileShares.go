package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeSmbfileSharesCmd = &cobra.Command{
	Use:   "describe-smbfile-shares",
	Short: "Gets a description for one or more Server Message Block (SMB) file shares from a S3 File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeSmbfileSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeSmbfileSharesCmd).Standalone()

		storagegateway_describeSmbfileSharesCmd.Flags().String("file-share-arnlist", "", "An array containing the Amazon Resource Name (ARN) of each file share to be described.")
		storagegateway_describeSmbfileSharesCmd.MarkFlagRequired("file-share-arnlist")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeSmbfileSharesCmd)
}
