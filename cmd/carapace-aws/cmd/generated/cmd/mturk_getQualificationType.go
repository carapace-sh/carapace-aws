package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getQualificationTypeCmd = &cobra.Command{
	Use:   "get-qualification-type",
	Short: "The `GetQualificationType`operation retrieves information about a Qualification type using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getQualificationTypeCmd).Standalone()

	mturk_getQualificationTypeCmd.Flags().String("qualification-type-id", "", "The ID of the QualificationType.")
	mturk_getQualificationTypeCmd.MarkFlagRequired("qualification-type-id")
	mturkCmd.AddCommand(mturk_getQualificationTypeCmd)
}
