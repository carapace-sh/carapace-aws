package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateRecipeCmd = &cobra.Command{
	Use:   "update-recipe",
	Short: "Modifies the definition of the `LATEST_WORKING` version of a DataBrew recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateRecipeCmd).Standalone()

	databrew_updateRecipeCmd.Flags().String("description", "", "A description of the recipe.")
	databrew_updateRecipeCmd.Flags().String("name", "", "The name of the recipe to be updated.")
	databrew_updateRecipeCmd.Flags().String("steps", "", "One or more steps to be performed by the recipe.")
	databrew_updateRecipeCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_updateRecipeCmd)
}
