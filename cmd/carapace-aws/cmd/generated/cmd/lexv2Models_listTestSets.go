package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listTestSetsCmd = &cobra.Command{
	Use:   "list-test-sets",
	Short: "The list of the test sets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listTestSetsCmd).Standalone()

	lexv2Models_listTestSetsCmd.Flags().String("max-results", "", "The maximum number of test sets to return in each page.")
	lexv2Models_listTestSetsCmd.Flags().String("next-token", "", "If the response from the ListTestSets operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
	lexv2Models_listTestSetsCmd.Flags().String("sort-by", "", "The sort order for the list of test sets.")
	lexv2ModelsCmd.AddCommand(lexv2Models_listTestSetsCmd)
}
