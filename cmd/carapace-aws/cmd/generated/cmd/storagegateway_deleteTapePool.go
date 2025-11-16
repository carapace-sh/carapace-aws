package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteTapePoolCmd = &cobra.Command{
	Use:   "delete-tape-pool",
	Short: "Delete a custom tape pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteTapePoolCmd).Standalone()

	storagegateway_deleteTapePoolCmd.Flags().String("pool-arn", "", "The Amazon Resource Name (ARN) of the custom tape pool to delete.")
	storagegateway_deleteTapePoolCmd.MarkFlagRequired("pool-arn")
	storagegatewayCmd.AddCommand(storagegateway_deleteTapePoolCmd)
}
