package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_createJobTemplateCmd = &cobra.Command{
	Use:   "create-job-template",
	Short: "Create a new job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_createJobTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_createJobTemplateCmd).Standalone()

		mediaconvert_createJobTemplateCmd.Flags().String("acceleration-settings", "", "Accelerated transcoding can significantly speed up jobs with long, visually complex content.")
		mediaconvert_createJobTemplateCmd.Flags().String("category", "", "Optional.")
		mediaconvert_createJobTemplateCmd.Flags().String("description", "", "Optional.")
		mediaconvert_createJobTemplateCmd.Flags().String("hop-destinations", "", "Optional.")
		mediaconvert_createJobTemplateCmd.Flags().String("name", "", "The name of the job template you are creating.")
		mediaconvert_createJobTemplateCmd.Flags().String("priority", "", "Specify the relative priority for this job.")
		mediaconvert_createJobTemplateCmd.Flags().String("queue", "", "Optional.")
		mediaconvert_createJobTemplateCmd.Flags().String("settings", "", "JobTemplateSettings contains all the transcode settings saved in the template that will be applied to jobs created from it.")
		mediaconvert_createJobTemplateCmd.Flags().String("status-update-interval", "", "Specify how often MediaConvert sends STATUS\\_UPDATE events to Amazon CloudWatch Events.")
		mediaconvert_createJobTemplateCmd.Flags().String("tags", "", "The tags that you want to add to the resource.")
		mediaconvert_createJobTemplateCmd.MarkFlagRequired("name")
		mediaconvert_createJobTemplateCmd.MarkFlagRequired("settings")
	})
	mediaconvertCmd.AddCommand(mediaconvert_createJobTemplateCmd)
}
