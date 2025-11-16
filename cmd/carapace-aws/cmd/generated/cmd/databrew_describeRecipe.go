package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeRecipeCmd = &cobra.Command{
	Use:   "describe-recipe",
	Short: "Returns the definition of a specific DataBrew recipe corresponding to a particular version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeRecipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_describeRecipeCmd).Standalone()

		databrew_describeRecipeCmd.Flags().String("name", "", "The name of the recipe to be described.")
		databrew_describeRecipeCmd.Flags().String("recipe-version", "", "The recipe version identifier.")
		databrew_describeRecipeCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_describeRecipeCmd)
}
