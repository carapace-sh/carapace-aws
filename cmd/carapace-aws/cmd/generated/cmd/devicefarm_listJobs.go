package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Gets information about jobs for a given test run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listJobsCmd).Standalone()

		devicefarm_listJobsCmd.Flags().String("arn", "", "The run's Amazon Resource Name (ARN).")
		devicefarm_listJobsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
		devicefarm_listJobsCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listJobsCmd)
}
