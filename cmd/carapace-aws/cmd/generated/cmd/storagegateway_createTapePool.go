package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createTapePoolCmd = &cobra.Command{
	Use:   "create-tape-pool",
	Short: "Creates a new custom tape pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createTapePoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createTapePoolCmd).Standalone()

		storagegateway_createTapePoolCmd.Flags().String("pool-name", "", "The name of the new custom tape pool.")
		storagegateway_createTapePoolCmd.Flags().String("retention-lock-time-in-days", "", "Tape retention lock time is set in days.")
		storagegateway_createTapePoolCmd.Flags().String("retention-lock-type", "", "Tape retention lock can be configured in two modes.")
		storagegateway_createTapePoolCmd.Flags().String("storage-class", "", "The storage class that is associated with the new custom pool.")
		storagegateway_createTapePoolCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to tape pool.")
		storagegateway_createTapePoolCmd.MarkFlagRequired("pool-name")
		storagegateway_createTapePoolCmd.MarkFlagRequired("storage-class")
	})
	storagegatewayCmd.AddCommand(storagegateway_createTapePoolCmd)
}
