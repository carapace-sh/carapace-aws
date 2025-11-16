package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listAnalysisTemplatesCmd = &cobra.Command{
	Use:   "list-analysis-templates",
	Short: "Lists analysis templates that the caller owns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listAnalysisTemplatesCmd).Standalone()

	cleanrooms_listAnalysisTemplatesCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listAnalysisTemplatesCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
	cleanrooms_listAnalysisTemplatesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listAnalysisTemplatesCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listAnalysisTemplatesCmd)
}
