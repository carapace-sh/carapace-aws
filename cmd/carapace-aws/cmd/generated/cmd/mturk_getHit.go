package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getHitCmd = &cobra.Command{
	Use:   "get-hit",
	Short: "The `GetHIT` operation retrieves the details of the specified HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getHitCmd).Standalone()

	mturk_getHitCmd.Flags().String("hitid", "", "The ID of the HIT to be retrieved.")
	mturk_getHitCmd.MarkFlagRequired("hitid")
	mturkCmd.AddCommand(mturk_getHitCmd)
}
