package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createRecipeCmd = &cobra.Command{
	Use:   "create-recipe",
	Short: "Creates a new DataBrew recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createRecipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_createRecipeCmd).Standalone()

		databrew_createRecipeCmd.Flags().String("description", "", "A description for the recipe.")
		databrew_createRecipeCmd.Flags().String("name", "", "A unique name for the recipe.")
		databrew_createRecipeCmd.Flags().String("steps", "", "An array containing the steps to be performed by the recipe.")
		databrew_createRecipeCmd.Flags().String("tags", "", "Metadata tags to apply to this recipe.")
		databrew_createRecipeCmd.MarkFlagRequired("name")
		databrew_createRecipeCmd.MarkFlagRequired("steps")
	})
	databrewCmd.AddCommand(databrew_createRecipeCmd)
}
