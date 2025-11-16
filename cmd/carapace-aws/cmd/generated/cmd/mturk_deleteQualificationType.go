package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_deleteQualificationTypeCmd = &cobra.Command{
	Use:   "delete-qualification-type",
	Short: "The `DeleteQualificationType` deletes a Qualification type and deletes any HIT types that are associated with the Qualification type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_deleteQualificationTypeCmd).Standalone()

	mturk_deleteQualificationTypeCmd.Flags().String("qualification-type-id", "", "The ID of the QualificationType to dispose.")
	mturk_deleteQualificationTypeCmd.MarkFlagRequired("qualification-type-id")
	mturkCmd.AddCommand(mturk_deleteQualificationTypeCmd)
}
