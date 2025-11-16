package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_startJobsQueryCmd = &cobra.Command{
	Use:   "start-jobs-query",
	Short: "Start an asynchronous jobs query using the provided filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_startJobsQueryCmd).Standalone()

	mediaconvert_startJobsQueryCmd.Flags().String("filter-list", "", "Optional.")
	mediaconvert_startJobsQueryCmd.Flags().String("max-results", "", "Optional.")
	mediaconvert_startJobsQueryCmd.Flags().String("next-token", "", "Use this string to request the next batch of jobs matched by a jobs query.")
	mediaconvert_startJobsQueryCmd.Flags().String("order", "", "Optional.")
	mediaconvertCmd.AddCommand(mediaconvert_startJobsQueryCmd)
}
