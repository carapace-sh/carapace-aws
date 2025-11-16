package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteTapeArchiveCmd = &cobra.Command{
	Use:   "delete-tape-archive",
	Short: "Deletes the specified virtual tape from the virtual tape shelf (VTS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteTapeArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteTapeArchiveCmd).Standalone()

		storagegateway_deleteTapeArchiveCmd.Flags().String("bypass-governance-retention", "", "Set to `TRUE` to delete an archived tape that belongs to a custom pool with tape retention lock.")
		storagegateway_deleteTapeArchiveCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape to delete from the virtual tape shelf (VTS).")
		storagegateway_deleteTapeArchiveCmd.MarkFlagRequired("tape-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteTapeArchiveCmd)
}
