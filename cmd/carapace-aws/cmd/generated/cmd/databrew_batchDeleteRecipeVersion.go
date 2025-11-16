package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_batchDeleteRecipeVersionCmd = &cobra.Command{
	Use:   "batch-delete-recipe-version",
	Short: "Deletes one or more versions of a recipe at a time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_batchDeleteRecipeVersionCmd).Standalone()

	databrew_batchDeleteRecipeVersionCmd.Flags().String("name", "", "The name of the recipe whose versions are to be deleted.")
	databrew_batchDeleteRecipeVersionCmd.Flags().String("recipe-versions", "", "An array of version identifiers, for the recipe versions to be deleted.")
	databrew_batchDeleteRecipeVersionCmd.MarkFlagRequired("name")
	databrew_batchDeleteRecipeVersionCmd.MarkFlagRequired("recipe-versions")
	databrewCmd.AddCommand(databrew_batchDeleteRecipeVersionCmd)
}
