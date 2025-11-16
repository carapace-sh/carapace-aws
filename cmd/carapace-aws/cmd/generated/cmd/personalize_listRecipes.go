package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listRecipesCmd = &cobra.Command{
	Use:   "list-recipes",
	Short: "Returns a list of available recipes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listRecipesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listRecipesCmd).Standalone()

		personalize_listRecipesCmd.Flags().String("domain", "", "Filters returned recipes by domain for a Domain dataset group.")
		personalize_listRecipesCmd.Flags().String("max-results", "", "The maximum number of recipes to return.")
		personalize_listRecipesCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListRecipes` for getting the next set of recipes (if they exist).")
		personalize_listRecipesCmd.Flags().String("recipe-provider", "", "The default is `SERVICE`.")
	})
	personalizeCmd.AddCommand(personalize_listRecipesCmd)
}
