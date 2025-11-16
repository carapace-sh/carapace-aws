package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUserDefinedFunctionCmd = &cobra.Command{
	Use:   "get-user-defined-function",
	Short: "Retrieves a specified function definition from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUserDefinedFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getUserDefinedFunctionCmd).Standalone()

		glue_getUserDefinedFunctionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the function to be retrieved is located.")
		glue_getUserDefinedFunctionCmd.Flags().String("database-name", "", "The name of the catalog database where the function is located.")
		glue_getUserDefinedFunctionCmd.Flags().String("function-name", "", "The name of the function.")
		glue_getUserDefinedFunctionCmd.MarkFlagRequired("database-name")
		glue_getUserDefinedFunctionCmd.MarkFlagRequired("function-name")
	})
	glueCmd.AddCommand(glue_getUserDefinedFunctionCmd)
}
