package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateDataCatalogCmd = &cobra.Command{
	Use:   "update-data-catalog",
	Short: "Updates the data catalog that has the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateDataCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_updateDataCatalogCmd).Standalone()

		athena_updateDataCatalogCmd.Flags().String("description", "", "New or modified text that describes the data catalog.")
		athena_updateDataCatalogCmd.Flags().String("name", "", "The name of the data catalog to update.")
		athena_updateDataCatalogCmd.Flags().String("parameters", "", "Specifies the Lambda function or functions to use for updating the data catalog.")
		athena_updateDataCatalogCmd.Flags().String("type", "", "Specifies the type of data catalog to update.")
		athena_updateDataCatalogCmd.MarkFlagRequired("name")
		athena_updateDataCatalogCmd.MarkFlagRequired("type")
	})
	athenaCmd.AddCommand(athena_updateDataCatalogCmd)
}
