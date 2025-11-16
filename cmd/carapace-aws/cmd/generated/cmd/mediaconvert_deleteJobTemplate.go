package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_deleteJobTemplateCmd = &cobra.Command{
	Use:   "delete-job-template",
	Short: "Permanently delete a job template you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_deleteJobTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_deleteJobTemplateCmd).Standalone()

		mediaconvert_deleteJobTemplateCmd.Flags().String("name", "", "The name of the job template to be deleted.")
		mediaconvert_deleteJobTemplateCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_deleteJobTemplateCmd)
}
