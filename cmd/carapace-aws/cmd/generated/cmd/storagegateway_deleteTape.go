package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteTapeCmd = &cobra.Command{
	Use:   "delete-tape",
	Short: "Deletes the specified virtual tape.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteTapeCmd).Standalone()

	storagegateway_deleteTapeCmd.Flags().String("bypass-governance-retention", "", "Set to `TRUE` to delete an archived tape that belongs to a custom pool with tape retention lock.")
	storagegateway_deleteTapeCmd.Flags().String("gateway-arn", "", "The unique Amazon Resource Name (ARN) of the gateway that the virtual tape to delete is associated with.")
	storagegateway_deleteTapeCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape to delete.")
	storagegateway_deleteTapeCmd.MarkFlagRequired("gateway-arn")
	storagegateway_deleteTapeCmd.MarkFlagRequired("tape-arn")
	storagegatewayCmd.AddCommand(storagegateway_deleteTapeCmd)
}
