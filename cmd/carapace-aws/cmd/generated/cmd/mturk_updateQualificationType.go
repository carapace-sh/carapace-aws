package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_updateQualificationTypeCmd = &cobra.Command{
	Use:   "update-qualification-type",
	Short: "The `UpdateQualificationType` operation modifies the attributes of an existing Qualification type, which is represented by a QualificationType data structure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_updateQualificationTypeCmd).Standalone()

	mturk_updateQualificationTypeCmd.Flags().String("answer-key", "", "The answers to the Qualification test specified in the Test parameter, in the form of an AnswerKey data structure.")
	mturk_updateQualificationTypeCmd.Flags().Bool("auto-granted", false, "Specifies whether requests for the Qualification type are granted immediately, without prompting the Worker with a Qualification test.")
	mturk_updateQualificationTypeCmd.Flags().String("auto-granted-value", "", "The Qualification value to use for automatically granted Qualifications.")
	mturk_updateQualificationTypeCmd.Flags().String("description", "", "The new description of the Qualification type.")
	mturk_updateQualificationTypeCmd.Flags().Bool("no-auto-granted", false, "Specifies whether requests for the Qualification type are granted immediately, without prompting the Worker with a Qualification test.")
	mturk_updateQualificationTypeCmd.Flags().String("qualification-type-id", "", "The ID of the Qualification type to update.")
	mturk_updateQualificationTypeCmd.Flags().String("qualification-type-status", "", "The new status of the Qualification type - Active | Inactive")
	mturk_updateQualificationTypeCmd.Flags().String("retry-delay-in-seconds", "", "The amount of time, in seconds, that Workers must wait after requesting a Qualification of the specified Qualification type before they can retry the Qualification request.")
	mturk_updateQualificationTypeCmd.Flags().String("test", "", "The questions for the Qualification test a Worker must answer correctly to obtain a Qualification of this type.")
	mturk_updateQualificationTypeCmd.Flags().String("test-duration-in-seconds", "", "The number of seconds the Worker has to complete the Qualification test, starting from the time the Worker requests the Qualification.")
	mturk_updateQualificationTypeCmd.Flag("no-auto-granted").Hidden = true
	mturk_updateQualificationTypeCmd.MarkFlagRequired("qualification-type-id")
	mturkCmd.AddCommand(mturk_updateQualificationTypeCmd)
}
