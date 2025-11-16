package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteRecipeVersionCmd = &cobra.Command{
	Use:   "delete-recipe-version",
	Short: "Deletes a single version of a DataBrew recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteRecipeVersionCmd).Standalone()

	databrew_deleteRecipeVersionCmd.Flags().String("name", "", "The name of the recipe.")
	databrew_deleteRecipeVersionCmd.Flags().String("recipe-version", "", "The version of the recipe to be deleted.")
	databrew_deleteRecipeVersionCmd.MarkFlagRequired("name")
	databrew_deleteRecipeVersionCmd.MarkFlagRequired("recipe-version")
	databrewCmd.AddCommand(databrew_deleteRecipeVersionCmd)
}
