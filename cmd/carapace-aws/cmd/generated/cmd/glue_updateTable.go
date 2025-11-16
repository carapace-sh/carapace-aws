package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateTableCmd = &cobra.Command{
	Use:   "update-table",
	Short: "Updates a metadata table in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateTableCmd).Standalone()

		glue_updateTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the table resides.")
		glue_updateTableCmd.Flags().String("database-name", "", "The name of the catalog database in which the table resides.")
		glue_updateTableCmd.Flags().Bool("force", false, "A flag that can be set to true to ignore matching storage descriptor and subobject matching requirements.")
		glue_updateTableCmd.Flags().String("name", "", "The unique identifier for the table within the specified database that will be created in the Glue Data Catalog.")
		glue_updateTableCmd.Flags().Bool("no-force", false, "A flag that can be set to true to ignore matching storage descriptor and subobject matching requirements.")
		glue_updateTableCmd.Flags().String("skip-archive", "", "By default, `UpdateTable` always creates an archived version of the table before updating it.")
		glue_updateTableCmd.Flags().String("table-input", "", "An updated `TableInput` object to define the metadata table in the catalog.")
		glue_updateTableCmd.Flags().String("transaction-id", "", "The transaction ID at which to update the table contents.")
		glue_updateTableCmd.Flags().String("update-open-table-format-input", "", "Input parameters for updating open table format tables in GlueData Catalog, serving as a wrapper for format-specific update operations such as Apache Iceberg.")
		glue_updateTableCmd.Flags().String("version-id", "", "The version ID at which to update the table contents.")
		glue_updateTableCmd.Flags().String("view-update-action", "", "The operation to be performed when updating the view.")
		glue_updateTableCmd.MarkFlagRequired("database-name")
		glue_updateTableCmd.Flag("no-force").Hidden = true
	})
	glueCmd.AddCommand(glue_updateTableCmd)
}
