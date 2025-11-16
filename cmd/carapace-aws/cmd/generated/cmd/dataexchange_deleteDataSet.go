package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_deleteDataSetCmd = &cobra.Command{
	Use:   "delete-data-set",
	Short: "This operation deletes a data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_deleteDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_deleteDataSetCmd).Standalone()

		dataexchange_deleteDataSetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_deleteDataSetCmd.MarkFlagRequired("data-set-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_deleteDataSetCmd)
}
