package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_publishRecipeCmd = &cobra.Command{
	Use:   "publish-recipe",
	Short: "Publishes a new version of a DataBrew recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_publishRecipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_publishRecipeCmd).Standalone()

		databrew_publishRecipeCmd.Flags().String("description", "", "A description of the recipe to be published, for this version of the recipe.")
		databrew_publishRecipeCmd.Flags().String("name", "", "The name of the recipe to be published.")
		databrew_publishRecipeCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_publishRecipeCmd)
}
