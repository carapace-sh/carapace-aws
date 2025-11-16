package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTableCmd = &cobra.Command{
	Use:   "get-table",
	Short: "Retrieves the `Table` definition in a Data Catalog for a specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTableCmd).Standalone()

	glue_getTableCmd.Flags().String("audit-context", "", "A structure containing the Lake Formation [audit context](https://docs.aws.amazon.com/glue/latest/webapi/API_AuditContext.html).")
	glue_getTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the table resides.")
	glue_getTableCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
	glue_getTableCmd.Flags().String("include-status-details", "", "Specifies whether to include status details related to a request to create or update an Glue Data Catalog view.")
	glue_getTableCmd.Flags().String("name", "", "The name of the table for which to retrieve the definition.")
	glue_getTableCmd.Flags().String("query-as-of-time", "", "The time as of when to read the table contents.")
	glue_getTableCmd.Flags().String("transaction-id", "", "The transaction ID at which to read the table contents.")
	glue_getTableCmd.MarkFlagRequired("database-name")
	glue_getTableCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_getTableCmd)
}
