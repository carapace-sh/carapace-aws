package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_updateHittypeOfHitCmd = &cobra.Command{
	Use:   "update-hittype-of-hit",
	Short: "The `UpdateHITTypeOfHIT` operation allows you to change the HITType properties of a HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_updateHittypeOfHitCmd).Standalone()

	mturk_updateHittypeOfHitCmd.Flags().String("hitid", "", "The HIT to update.")
	mturk_updateHittypeOfHitCmd.Flags().String("hittype-id", "", "The ID of the new HIT type.")
	mturk_updateHittypeOfHitCmd.MarkFlagRequired("hitid")
	mturk_updateHittypeOfHitCmd.MarkFlagRequired("hittype-id")
	mturkCmd.AddCommand(mturk_updateHittypeOfHitCmd)
}
