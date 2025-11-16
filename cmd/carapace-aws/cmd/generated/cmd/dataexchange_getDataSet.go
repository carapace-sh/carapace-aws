package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getDataSetCmd = &cobra.Command{
	Use:   "get-data-set",
	Short: "This operation returns information about a data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_getDataSetCmd).Standalone()

		dataexchange_getDataSetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_getDataSetCmd.MarkFlagRequired("data-set-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_getDataSetCmd)
}
