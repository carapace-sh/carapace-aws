package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_createQualificationTypeCmd = &cobra.Command{
	Use:   "create-qualification-type",
	Short: "The `CreateQualificationType` operation creates a new Qualification type, which is represented by a `QualificationType` data structure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_createQualificationTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_createQualificationTypeCmd).Standalone()

		mturk_createQualificationTypeCmd.Flags().String("answer-key", "", "The answers to the Qualification test specified in the Test parameter, in the form of an AnswerKey data structure.")
		mturk_createQualificationTypeCmd.Flags().Bool("auto-granted", false, "Specifies whether requests for the Qualification type are granted immediately, without prompting the Worker with a Qualification test.")
		mturk_createQualificationTypeCmd.Flags().String("auto-granted-value", "", "The Qualification value to use for automatically granted Qualifications.")
		mturk_createQualificationTypeCmd.Flags().String("description", "", "A long description for the Qualification type.")
		mturk_createQualificationTypeCmd.Flags().String("keywords", "", "One or more words or phrases that describe the Qualification type, separated by commas.")
		mturk_createQualificationTypeCmd.Flags().String("name", "", "The name you give to the Qualification type.")
		mturk_createQualificationTypeCmd.Flags().Bool("no-auto-granted", false, "Specifies whether requests for the Qualification type are granted immediately, without prompting the Worker with a Qualification test.")
		mturk_createQualificationTypeCmd.Flags().String("qualification-type-status", "", "The initial status of the Qualification type.")
		mturk_createQualificationTypeCmd.Flags().String("retry-delay-in-seconds", "", "The number of seconds that a Worker must wait after requesting a Qualification of the Qualification type before the worker can retry the Qualification request.")
		mturk_createQualificationTypeCmd.Flags().String("test", "", "The questions for the Qualification test a Worker must answer correctly to obtain a Qualification of this type.")
		mturk_createQualificationTypeCmd.Flags().String("test-duration-in-seconds", "", "The number of seconds the Worker has to complete the Qualification test, starting from the time the Worker requests the Qualification.")
		mturk_createQualificationTypeCmd.MarkFlagRequired("description")
		mturk_createQualificationTypeCmd.MarkFlagRequired("name")
		mturk_createQualificationTypeCmd.Flag("no-auto-granted").Hidden = true
		mturk_createQualificationTypeCmd.MarkFlagRequired("qualification-type-status")
	})
	mturkCmd.AddCommand(mturk_createQualificationTypeCmd)
}
