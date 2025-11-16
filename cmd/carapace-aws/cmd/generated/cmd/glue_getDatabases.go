package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDatabasesCmd = &cobra.Command{
	Use:   "get-databases",
	Short: "Retrieves all databases defined in a given Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getDatabasesCmd).Standalone()

		glue_getDatabasesCmd.Flags().String("attributes-to-get", "", "Specifies the database fields returned by the `GetDatabases` call.")
		glue_getDatabasesCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog from which to retrieve `Databases`.")
		glue_getDatabasesCmd.Flags().String("max-results", "", "The maximum number of databases to return in one response.")
		glue_getDatabasesCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_getDatabasesCmd.Flags().String("resource-share-type", "", "Allows you to specify that you want to list the databases shared with your account.")
	})
	glueCmd.AddCommand(glue_getDatabasesCmd)
}
