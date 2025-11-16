package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_listEntitiesDetectionV2JobsCmd = &cobra.Command{
	Use:   "list-entities-detection-v2-jobs",
	Short: "Gets a list of medical entity detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_listEntitiesDetectionV2JobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_listEntitiesDetectionV2JobsCmd).Standalone()

		comprehendmedical_listEntitiesDetectionV2JobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
		comprehendmedical_listEntitiesDetectionV2JobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
		comprehendmedical_listEntitiesDetectionV2JobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_listEntitiesDetectionV2JobsCmd)
}
