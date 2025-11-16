package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteUserDefinedFunctionCmd = &cobra.Command{
	Use:   "delete-user-defined-function",
	Short: "Deletes an existing function definition from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteUserDefinedFunctionCmd).Standalone()

	glue_deleteUserDefinedFunctionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the function to be deleted is located.")
	glue_deleteUserDefinedFunctionCmd.Flags().String("database-name", "", "The name of the catalog database where the function is located.")
	glue_deleteUserDefinedFunctionCmd.Flags().String("function-name", "", "The name of the function definition to be deleted.")
	glue_deleteUserDefinedFunctionCmd.MarkFlagRequired("database-name")
	glue_deleteUserDefinedFunctionCmd.MarkFlagRequired("function-name")
	glueCmd.AddCommand(glue_deleteUserDefinedFunctionCmd)
}
