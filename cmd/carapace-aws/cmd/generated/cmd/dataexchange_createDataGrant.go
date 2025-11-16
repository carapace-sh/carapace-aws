package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_createDataGrantCmd = &cobra.Command{
	Use:   "create-data-grant",
	Short: "This operation creates a data grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_createDataGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_createDataGrantCmd).Standalone()

		dataexchange_createDataGrantCmd.Flags().String("description", "", "The description of the data grant.")
		dataexchange_createDataGrantCmd.Flags().String("ends-at", "", "The timestamp of when access to the associated data set ends.")
		dataexchange_createDataGrantCmd.Flags().String("grant-distribution-scope", "", "The distribution scope of the data grant.")
		dataexchange_createDataGrantCmd.Flags().String("name", "", "The name of the data grant.")
		dataexchange_createDataGrantCmd.Flags().String("receiver-principal", "", "The Amazon Web Services account ID of the data grant receiver.")
		dataexchange_createDataGrantCmd.Flags().String("source-data-set-id", "", "The ID of the data set used to create the data grant.")
		dataexchange_createDataGrantCmd.Flags().String("tags", "", "The tags to add to the data grant.")
		dataexchange_createDataGrantCmd.MarkFlagRequired("grant-distribution-scope")
		dataexchange_createDataGrantCmd.MarkFlagRequired("name")
		dataexchange_createDataGrantCmd.MarkFlagRequired("receiver-principal")
		dataexchange_createDataGrantCmd.MarkFlagRequired("source-data-set-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_createDataGrantCmd)
}
