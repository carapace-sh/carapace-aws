package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_assignTapePoolCmd = &cobra.Command{
	Use:   "assign-tape-pool",
	Short: "Assigns a tape to a tape pool for archiving.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_assignTapePoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_assignTapePoolCmd).Standalone()

		storagegateway_assignTapePoolCmd.Flags().String("bypass-governance-retention", "", "Set permissions to bypass governance retention.")
		storagegateway_assignTapePoolCmd.Flags().String("pool-id", "", "The ID of the pool that you want to add your tape to for archiving.")
		storagegateway_assignTapePoolCmd.Flags().String("tape-arn", "", "The unique Amazon Resource Name (ARN) of the virtual tape that you want to add to the tape pool.")
		storagegateway_assignTapePoolCmd.MarkFlagRequired("pool-id")
		storagegateway_assignTapePoolCmd.MarkFlagRequired("tape-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_assignTapePoolCmd)
}
