package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUserDefinedFunctionsCmd = &cobra.Command{
	Use:   "get-user-defined-functions",
	Short: "Retrieves multiple function definitions from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUserDefinedFunctionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getUserDefinedFunctionsCmd).Standalone()

		glue_getUserDefinedFunctionsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the functions to be retrieved are located.")
		glue_getUserDefinedFunctionsCmd.Flags().String("database-name", "", "The name of the catalog database where the functions are located.")
		glue_getUserDefinedFunctionsCmd.Flags().String("max-results", "", "The maximum number of functions to return in one response.")
		glue_getUserDefinedFunctionsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_getUserDefinedFunctionsCmd.Flags().String("pattern", "", "An optional function-name pattern string that filters the function definitions returned.")
		glue_getUserDefinedFunctionsCmd.MarkFlagRequired("pattern")
	})
	glueCmd.AddCommand(glue_getUserDefinedFunctionsCmd)
}
