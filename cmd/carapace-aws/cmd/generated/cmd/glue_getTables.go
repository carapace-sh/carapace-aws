package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTablesCmd = &cobra.Command{
	Use:   "get-tables",
	Short: "Retrieves the definitions of some or all of the tables in a given `Database`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTablesCmd).Standalone()

	glue_getTablesCmd.Flags().String("attributes-to-get", "", "Specifies the table fields returned by the `GetTables` call.")
	glue_getTablesCmd.Flags().String("audit-context", "", "A structure containing the Lake Formation [audit context](https://docs.aws.amazon.com/glue/latest/webapi/API_AuditContext.html).")
	glue_getTablesCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the tables reside.")
	glue_getTablesCmd.Flags().String("database-name", "", "The database in the catalog whose tables to list.")
	glue_getTablesCmd.Flags().String("expression", "", "A regular expression pattern.")
	glue_getTablesCmd.Flags().String("include-status-details", "", "Specifies whether to include status details related to a request to create or update an Glue Data Catalog view.")
	glue_getTablesCmd.Flags().String("max-results", "", "The maximum number of tables to return in a single response.")
	glue_getTablesCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
	glue_getTablesCmd.Flags().String("query-as-of-time", "", "The time as of when to read the table contents.")
	glue_getTablesCmd.Flags().String("transaction-id", "", "The transaction ID at which to read the table contents.")
	glue_getTablesCmd.MarkFlagRequired("database-name")
	glueCmd.AddCommand(glue_getTablesCmd)
}
