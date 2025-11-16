package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getReceivedDataGrantCmd = &cobra.Command{
	Use:   "get-received-data-grant",
	Short: "This operation returns information about a received data grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getReceivedDataGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_getReceivedDataGrantCmd).Standalone()

		dataexchange_getReceivedDataGrantCmd.Flags().String("data-grant-arn", "", "The Amazon Resource Name (ARN) of the data grant.")
		dataexchange_getReceivedDataGrantCmd.MarkFlagRequired("data-grant-arn")
	})
	dataexchangeCmd.AddCommand(dataexchange_getReceivedDataGrantCmd)
}
