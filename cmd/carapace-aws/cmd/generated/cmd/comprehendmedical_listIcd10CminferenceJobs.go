package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_listIcd10CminferenceJobsCmd = &cobra.Command{
	Use:   "list-icd10-cminference-jobs",
	Short: "Gets a list of InferICD10CM jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_listIcd10CminferenceJobsCmd).Standalone()

	comprehendmedical_listIcd10CminferenceJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehendmedical_listIcd10CminferenceJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehendmedical_listIcd10CminferenceJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendmedicalCmd.AddCommand(comprehendmedical_listIcd10CminferenceJobsCmd)
}
