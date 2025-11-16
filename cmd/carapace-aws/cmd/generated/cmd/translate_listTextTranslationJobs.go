package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_listTextTranslationJobsCmd = &cobra.Command{
	Use:   "list-text-translation-jobs",
	Short: "Gets a list of the batch translation jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_listTextTranslationJobsCmd).Standalone()

	translate_listTextTranslationJobsCmd.Flags().String("filter", "", "The parameters that specify which batch translation jobs to retrieve.")
	translate_listTextTranslationJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	translate_listTextTranslationJobsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	translateCmd.AddCommand(translate_listTextTranslationJobsCmd)
}
