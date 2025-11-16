package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listRecipesCmd = &cobra.Command{
	Use:   "list-recipes",
	Short: "Lists all of the DataBrew recipes that are defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listRecipesCmd).Standalone()

	databrew_listRecipesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	databrew_listRecipesCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	databrew_listRecipesCmd.Flags().String("recipe-version", "", "Return only those recipes with a version identifier of `LATEST_WORKING` or `LATEST_PUBLISHED`.")
	databrewCmd.AddCommand(databrew_listRecipesCmd)
}
