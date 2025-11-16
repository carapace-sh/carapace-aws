package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listRecipeVersionsCmd = &cobra.Command{
	Use:   "list-recipe-versions",
	Short: "Lists the versions of a particular DataBrew recipe, except for `LATEST_WORKING`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listRecipeVersionsCmd).Standalone()

	databrew_listRecipeVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	databrew_listRecipeVersionsCmd.Flags().String("name", "", "The name of the recipe for which to return version information.")
	databrew_listRecipeVersionsCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	databrew_listRecipeVersionsCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_listRecipeVersionsCmd)
}
