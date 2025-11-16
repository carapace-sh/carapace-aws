package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_copyJobTemplateCmd = &cobra.Command{
	Use:   "copy-job-template",
	Short: "Copies a job template to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_copyJobTemplateCmd).Standalone()

	deadline_copyJobTemplateCmd.Flags().String("farm-id", "", "The farm ID to copy.")
	deadline_copyJobTemplateCmd.Flags().String("job-id", "", "The job ID to copy.")
	deadline_copyJobTemplateCmd.Flags().String("queue-id", "", "The queue ID to copy.")
	deadline_copyJobTemplateCmd.Flags().String("target-s3-location", "", "The Amazon S3 bucket name and key where you would like to add a copy of the job template.")
	deadline_copyJobTemplateCmd.MarkFlagRequired("farm-id")
	deadline_copyJobTemplateCmd.MarkFlagRequired("job-id")
	deadline_copyJobTemplateCmd.MarkFlagRequired("queue-id")
	deadline_copyJobTemplateCmd.MarkFlagRequired("target-s3-location")
	deadlineCmd.AddCommand(deadline_copyJobTemplateCmd)
}
