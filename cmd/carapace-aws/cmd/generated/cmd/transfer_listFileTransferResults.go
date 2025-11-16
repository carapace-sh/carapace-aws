package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listFileTransferResultsCmd = &cobra.Command{
	Use:   "list-file-transfer-results",
	Short: "Returns real-time updates and detailed information on the status of each individual file being transferred in a specific file transfer operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listFileTransferResultsCmd).Standalone()

	transfer_listFileTransferResultsCmd.Flags().String("connector-id", "", "A unique identifier for a connector.")
	transfer_listFileTransferResultsCmd.Flags().String("max-results", "", "The maximum number of files to return in a single page.")
	transfer_listFileTransferResultsCmd.Flags().String("next-token", "", "If there are more file details than returned in this call, use this value for a subsequent call to `ListFileTransferResults` to retrieve them.")
	transfer_listFileTransferResultsCmd.Flags().String("transfer-id", "", "A unique identifier for a file transfer.")
	transfer_listFileTransferResultsCmd.MarkFlagRequired("connector-id")
	transfer_listFileTransferResultsCmd.MarkFlagRequired("transfer-id")
	transferCmd.AddCommand(transfer_listFileTransferResultsCmd)
}
