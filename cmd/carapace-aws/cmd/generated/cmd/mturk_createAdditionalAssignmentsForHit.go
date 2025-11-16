package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_createAdditionalAssignmentsForHitCmd = &cobra.Command{
	Use:   "create-additional-assignments-for-hit",
	Short: "The `CreateAdditionalAssignmentsForHIT` operation increases the maximum number of assignments of an existing HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_createAdditionalAssignmentsForHitCmd).Standalone()

	mturk_createAdditionalAssignmentsForHitCmd.Flags().String("hitid", "", "The ID of the HIT to extend.")
	mturk_createAdditionalAssignmentsForHitCmd.Flags().String("number-of-additional-assignments", "", "The number of additional assignments to request for this HIT.")
	mturk_createAdditionalAssignmentsForHitCmd.Flags().String("unique-request-token", "", "A unique identifier for this request, which allows you to retry the call on error without extending the HIT multiple times.")
	mturk_createAdditionalAssignmentsForHitCmd.MarkFlagRequired("hitid")
	mturk_createAdditionalAssignmentsForHitCmd.MarkFlagRequired("number-of-additional-assignments")
	mturkCmd.AddCommand(mturk_createAdditionalAssignmentsForHitCmd)
}
