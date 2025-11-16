package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_updateDataSetCmd = &cobra.Command{
	Use:   "update-data-set",
	Short: "This operation updates a data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_updateDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_updateDataSetCmd).Standalone()

		dataexchange_updateDataSetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_updateDataSetCmd.Flags().String("description", "", "The description for the data set.")
		dataexchange_updateDataSetCmd.Flags().String("name", "", "The name of the data set.")
		dataexchange_updateDataSetCmd.MarkFlagRequired("data-set-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_updateDataSetCmd)
}
