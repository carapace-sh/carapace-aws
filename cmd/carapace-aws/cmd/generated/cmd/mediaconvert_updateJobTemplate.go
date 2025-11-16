package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_updateJobTemplateCmd = &cobra.Command{
	Use:   "update-job-template",
	Short: "Modify one of your existing job templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_updateJobTemplateCmd).Standalone()

	mediaconvert_updateJobTemplateCmd.Flags().String("acceleration-settings", "", "Accelerated transcoding can significantly speed up jobs with long, visually complex content.")
	mediaconvert_updateJobTemplateCmd.Flags().String("category", "", "The new category for the job template, if you are changing it.")
	mediaconvert_updateJobTemplateCmd.Flags().String("description", "", "The new description for the job template, if you are changing it.")
	mediaconvert_updateJobTemplateCmd.Flags().String("hop-destinations", "", "Optional list of hop destinations.")
	mediaconvert_updateJobTemplateCmd.Flags().String("name", "", "The name of the job template you are modifying")
	mediaconvert_updateJobTemplateCmd.Flags().String("priority", "", "Specify the relative priority for this job.")
	mediaconvert_updateJobTemplateCmd.Flags().String("queue", "", "The new queue for the job template, if you are changing it.")
	mediaconvert_updateJobTemplateCmd.Flags().String("settings", "", "JobTemplateSettings contains all the transcode settings saved in the template that will be applied to jobs created from it.")
	mediaconvert_updateJobTemplateCmd.Flags().String("status-update-interval", "", "Specify how often MediaConvert sends STATUS\\_UPDATE events to Amazon CloudWatch Events.")
	mediaconvert_updateJobTemplateCmd.MarkFlagRequired("name")
	mediaconvertCmd.AddCommand(mediaconvert_updateJobTemplateCmd)
}
