package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_listSnomedctinferenceJobsCmd = &cobra.Command{
	Use:   "list-snomedctinference-jobs",
	Short: "Gets a list of InferSNOMEDCT jobs a user has submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_listSnomedctinferenceJobsCmd).Standalone()

	comprehendmedical_listSnomedctinferenceJobsCmd.Flags().String("filter", "", "")
	comprehendmedical_listSnomedctinferenceJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehendmedical_listSnomedctinferenceJobsCmd.Flags().String("next-token", "", "Identifies the next page of InferSNOMEDCT results to return.")
	comprehendmedicalCmd.AddCommand(comprehendmedical_listSnomedctinferenceJobsCmd)
}
