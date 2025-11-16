package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_disassociateQualificationFromWorkerCmd = &cobra.Command{
	Use:   "disassociate-qualification-from-worker",
	Short: "The `DisassociateQualificationFromWorker` revokes a previously granted Qualification from a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_disassociateQualificationFromWorkerCmd).Standalone()

	mturk_disassociateQualificationFromWorkerCmd.Flags().String("qualification-type-id", "", "The ID of the Qualification type of the Qualification to be revoked.")
	mturk_disassociateQualificationFromWorkerCmd.Flags().String("reason", "", "A text message that explains why the Qualification was revoked.")
	mturk_disassociateQualificationFromWorkerCmd.Flags().String("worker-id", "", "The ID of the Worker who possesses the Qualification to be revoked.")
	mturk_disassociateQualificationFromWorkerCmd.MarkFlagRequired("qualification-type-id")
	mturk_disassociateQualificationFromWorkerCmd.MarkFlagRequired("worker-id")
	mturkCmd.AddCommand(mturk_disassociateQualificationFromWorkerCmd)
}
