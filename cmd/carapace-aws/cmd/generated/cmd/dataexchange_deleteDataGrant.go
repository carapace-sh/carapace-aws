package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_deleteDataGrantCmd = &cobra.Command{
	Use:   "delete-data-grant",
	Short: "This operation deletes a data grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_deleteDataGrantCmd).Standalone()

	dataexchange_deleteDataGrantCmd.Flags().String("data-grant-id", "", "The ID of the data grant to delete.")
	dataexchange_deleteDataGrantCmd.MarkFlagRequired("data-grant-id")
	dataexchangeCmd.AddCommand(dataexchange_deleteDataGrantCmd)
}
