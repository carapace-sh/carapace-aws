package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_deleteHitCmd = &cobra.Command{
	Use:   "delete-hit",
	Short: "The `DeleteHIT` operation is used to delete HIT that is no longer needed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_deleteHitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_deleteHitCmd).Standalone()

		mturk_deleteHitCmd.Flags().String("hitid", "", "The ID of the HIT to be deleted.")
		mturk_deleteHitCmd.MarkFlagRequired("hitid")
	})
	mturkCmd.AddCommand(mturk_deleteHitCmd)
}
