package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getJobBookmarkCmd = &cobra.Command{
	Use:   "get-job-bookmark",
	Short: "Returns information on a job bookmark entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getJobBookmarkCmd).Standalone()

	glue_getJobBookmarkCmd.Flags().String("job-name", "", "The name of the job in question.")
	glue_getJobBookmarkCmd.Flags().String("run-id", "", "The unique run identifier associated with this job run.")
	glue_getJobBookmarkCmd.MarkFlagRequired("job-name")
	glueCmd.AddCommand(glue_getJobBookmarkCmd)
}
