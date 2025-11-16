package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateUserDefinedFunctionCmd = &cobra.Command{
	Use:   "update-user-defined-function",
	Short: "Updates an existing function definition in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateUserDefinedFunctionCmd).Standalone()

	glue_updateUserDefinedFunctionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the function to be updated is located.")
	glue_updateUserDefinedFunctionCmd.Flags().String("database-name", "", "The name of the catalog database where the function to be updated is located.")
	glue_updateUserDefinedFunctionCmd.Flags().String("function-input", "", "A `FunctionInput` object that redefines the function in the Data Catalog.")
	glue_updateUserDefinedFunctionCmd.Flags().String("function-name", "", "The name of the function.")
	glue_updateUserDefinedFunctionCmd.MarkFlagRequired("database-name")
	glue_updateUserDefinedFunctionCmd.MarkFlagRequired("function-input")
	glue_updateUserDefinedFunctionCmd.MarkFlagRequired("function-name")
	glueCmd.AddCommand(glue_updateUserDefinedFunctionCmd)
}
