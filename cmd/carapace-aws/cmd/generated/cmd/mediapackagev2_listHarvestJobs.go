package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_listHarvestJobsCmd = &cobra.Command{
	Use:   "list-harvest-jobs",
	Short: "Retrieves a list of harvest jobs that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_listHarvestJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_listHarvestJobsCmd).Standalone()

		mediapackagev2_listHarvestJobsCmd.Flags().String("channel-group-name", "", "The name of the channel group to filter the harvest jobs by.")
		mediapackagev2_listHarvestJobsCmd.Flags().String("channel-name", "", "The name of the channel to filter the harvest jobs by.")
		mediapackagev2_listHarvestJobsCmd.Flags().String("max-results", "", "The maximum number of harvest jobs to return in a single request.")
		mediapackagev2_listHarvestJobsCmd.Flags().String("next-token", "", "A token used for pagination.")
		mediapackagev2_listHarvestJobsCmd.Flags().String("origin-endpoint-name", "", "The name of the origin endpoint to filter the harvest jobs by.")
		mediapackagev2_listHarvestJobsCmd.Flags().String("status", "", "The status to filter the harvest jobs by.")
		mediapackagev2_listHarvestJobsCmd.MarkFlagRequired("channel-group-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_listHarvestJobsCmd)
}
