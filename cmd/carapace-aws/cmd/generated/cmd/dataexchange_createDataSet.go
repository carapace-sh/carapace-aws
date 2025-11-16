package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_createDataSetCmd = &cobra.Command{
	Use:   "create-data-set",
	Short: "This operation creates a data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_createDataSetCmd).Standalone()

	dataexchange_createDataSetCmd.Flags().String("asset-type", "", "The type of asset that is added to a data set.")
	dataexchange_createDataSetCmd.Flags().String("description", "", "A description for the data set.")
	dataexchange_createDataSetCmd.Flags().String("name", "", "The name of the data set.")
	dataexchange_createDataSetCmd.Flags().String("tags", "", "A data set tag is an optional label that you can assign to a data set when you create it.")
	dataexchange_createDataSetCmd.MarkFlagRequired("asset-type")
	dataexchange_createDataSetCmd.MarkFlagRequired("description")
	dataexchange_createDataSetCmd.MarkFlagRequired("name")
	dataexchangeCmd.AddCommand(dataexchange_createDataSetCmd)
}
