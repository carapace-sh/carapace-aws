package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_resetJobBookmarkCmd = &cobra.Command{
	Use:   "reset-job-bookmark",
	Short: "Resets a bookmark entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_resetJobBookmarkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_resetJobBookmarkCmd).Standalone()

		glue_resetJobBookmarkCmd.Flags().String("job-name", "", "The name of the job in question.")
		glue_resetJobBookmarkCmd.Flags().String("run-id", "", "The unique run identifier associated with this job run.")
		glue_resetJobBookmarkCmd.MarkFlagRequired("job-name")
	})
	glueCmd.AddCommand(glue_resetJobBookmarkCmd)
}
