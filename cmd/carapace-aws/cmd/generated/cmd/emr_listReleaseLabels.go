package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listReleaseLabelsCmd = &cobra.Command{
	Use:   "list-release-labels",
	Short: "Retrieves release labels of Amazon EMR services in the Region where the API is called.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listReleaseLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listReleaseLabelsCmd).Standalone()

		emr_listReleaseLabelsCmd.Flags().String("filters", "", "Filters the results of the request.")
		emr_listReleaseLabelsCmd.Flags().String("max-results", "", "Defines the maximum number of release labels to return in a single response.")
		emr_listReleaseLabelsCmd.Flags().String("next-token", "", "Specifies the next page of results.")
	})
	emrCmd.AddCommand(emr_listReleaseLabelsCmd)
}
