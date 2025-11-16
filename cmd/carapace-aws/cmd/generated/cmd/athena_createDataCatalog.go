package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createDataCatalogCmd = &cobra.Command{
	Use:   "create-data-catalog",
	Short: "Creates (registers) a data catalog with the specified name and properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createDataCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_createDataCatalogCmd).Standalone()

		athena_createDataCatalogCmd.Flags().String("description", "", "A description of the data catalog to be created.")
		athena_createDataCatalogCmd.Flags().String("name", "", "The name of the data catalog to create.")
		athena_createDataCatalogCmd.Flags().String("parameters", "", "Specifies the Lambda function or functions to use for creating the data catalog.")
		athena_createDataCatalogCmd.Flags().String("tags", "", "A list of comma separated tags to add to the data catalog that is created.")
		athena_createDataCatalogCmd.Flags().String("type", "", "The type of data catalog to create: `LAMBDA` for a federated catalog, `GLUE` for an Glue Data Catalog, and `HIVE` for an external Apache Hive metastore.")
		athena_createDataCatalogCmd.MarkFlagRequired("name")
		athena_createDataCatalogCmd.MarkFlagRequired("type")
	})
	athenaCmd.AddCommand(athena_createDataCatalogCmd)
}
