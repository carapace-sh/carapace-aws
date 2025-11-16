package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_updateHitreviewStatusCmd = &cobra.Command{
	Use:   "update-hitreview-status",
	Short: "The `UpdateHITReviewStatus` operation updates the status of a HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_updateHitreviewStatusCmd).Standalone()

	mturk_updateHitreviewStatusCmd.Flags().String("hitid", "", "The ID of the HIT to update.")
	mturk_updateHitreviewStatusCmd.Flags().Bool("no-revert", false, "Specifies how to update the HIT status.")
	mturk_updateHitreviewStatusCmd.Flags().Bool("revert", false, "Specifies how to update the HIT status.")
	mturk_updateHitreviewStatusCmd.MarkFlagRequired("hitid")
	mturk_updateHitreviewStatusCmd.Flag("no-revert").Hidden = true
	mturkCmd.AddCommand(mturk_updateHitreviewStatusCmd)
}
