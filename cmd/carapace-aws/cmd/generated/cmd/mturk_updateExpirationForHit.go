package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_updateExpirationForHitCmd = &cobra.Command{
	Use:   "update-expiration-for-hit",
	Short: "The `UpdateExpirationForHIT` operation allows you update the expiration time of a HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_updateExpirationForHitCmd).Standalone()

	mturk_updateExpirationForHitCmd.Flags().String("expire-at", "", "The date and time at which you want the HIT to expire")
	mturk_updateExpirationForHitCmd.Flags().String("hitid", "", "The HIT to update.")
	mturk_updateExpirationForHitCmd.MarkFlagRequired("expire-at")
	mturk_updateExpirationForHitCmd.MarkFlagRequired("hitid")
	mturkCmd.AddCommand(mturk_updateExpirationForHitCmd)
}
