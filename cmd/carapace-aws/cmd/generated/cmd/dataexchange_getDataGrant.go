package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getDataGrantCmd = &cobra.Command{
	Use:   "get-data-grant",
	Short: "This operation returns information about a data grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getDataGrantCmd).Standalone()

	dataexchange_getDataGrantCmd.Flags().String("data-grant-id", "", "The ID of the data grant.")
	dataexchange_getDataGrantCmd.MarkFlagRequired("data-grant-id")
	dataexchangeCmd.AddCommand(dataexchange_getDataGrantCmd)
}
