package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_deletePlaybackRestrictionPolicyCmd = &cobra.Command{
	Use:   "delete-playback-restriction-policy",
	Short: "Deletes the specified playback restriction policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_deletePlaybackRestrictionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_deletePlaybackRestrictionPolicyCmd).Standalone()

		ivs_deletePlaybackRestrictionPolicyCmd.Flags().String("arn", "", "ARN of the playback restriction policy to be deleted.")
		ivs_deletePlaybackRestrictionPolicyCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_deletePlaybackRestrictionPolicyCmd)
}
