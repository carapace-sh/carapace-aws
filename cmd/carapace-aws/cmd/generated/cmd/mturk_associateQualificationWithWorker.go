package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_associateQualificationWithWorkerCmd = &cobra.Command{
	Use:   "associate-qualification-with-worker",
	Short: "The `AssociateQualificationWithWorker` operation gives a Worker a Qualification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_associateQualificationWithWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_associateQualificationWithWorkerCmd).Standalone()

		mturk_associateQualificationWithWorkerCmd.Flags().String("integer-value", "", "The value of the Qualification to assign.")
		mturk_associateQualificationWithWorkerCmd.Flags().Bool("no-send-notification", false, "Specifies whether to send a notification email message to the Worker saying that the qualification was assigned to the Worker.")
		mturk_associateQualificationWithWorkerCmd.Flags().String("qualification-type-id", "", "The ID of the Qualification type to use for the assigned Qualification.")
		mturk_associateQualificationWithWorkerCmd.Flags().Bool("send-notification", false, "Specifies whether to send a notification email message to the Worker saying that the qualification was assigned to the Worker.")
		mturk_associateQualificationWithWorkerCmd.Flags().String("worker-id", "", "The ID of the Worker to whom the Qualification is being assigned.")
		mturk_associateQualificationWithWorkerCmd.Flag("no-send-notification").Hidden = true
		mturk_associateQualificationWithWorkerCmd.MarkFlagRequired("qualification-type-id")
		mturk_associateQualificationWithWorkerCmd.MarkFlagRequired("worker-id")
	})
	mturkCmd.AddCommand(mturk_associateQualificationWithWorkerCmd)
}
