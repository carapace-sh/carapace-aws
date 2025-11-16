package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createUserDefinedFunctionCmd = &cobra.Command{
	Use:   "create-user-defined-function",
	Short: "Creates a new function definition in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createUserDefinedFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createUserDefinedFunctionCmd).Standalone()

		glue_createUserDefinedFunctionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which to create the function.")
		glue_createUserDefinedFunctionCmd.Flags().String("database-name", "", "The name of the catalog database in which to create the function.")
		glue_createUserDefinedFunctionCmd.Flags().String("function-input", "", "A `FunctionInput` object that defines the function to create in the Data Catalog.")
		glue_createUserDefinedFunctionCmd.MarkFlagRequired("database-name")
		glue_createUserDefinedFunctionCmd.MarkFlagRequired("function-input")
	})
	glueCmd.AddCommand(glue_createUserDefinedFunctionCmd)
}
