package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listFileSharesCmd = &cobra.Command{
	Use:   "list-file-shares",
	Short: "Gets a list of the file shares for a specific S3 File Gateway, or the list of file shares that belong to the calling Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listFileSharesCmd).Standalone()

	storagegateway_listFileSharesCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway whose file shares you want to list.")
	storagegateway_listFileSharesCmd.Flags().String("limit", "", "The maximum number of file shares to return in the response.")
	storagegateway_listFileSharesCmd.Flags().String("marker", "", "Opaque pagination token returned from a previous ListFileShares operation.")
	storagegatewayCmd.AddCommand(storagegateway_listFileSharesCmd)
}
