package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getPlaybackRestrictionPolicyCmd = &cobra.Command{
	Use:   "get-playback-restriction-policy",
	Short: "Gets the specified playback restriction policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getPlaybackRestrictionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_getPlaybackRestrictionPolicyCmd).Standalone()

		ivs_getPlaybackRestrictionPolicyCmd.Flags().String("arn", "", "ARN of the playback restriction policy to be returned.")
		ivs_getPlaybackRestrictionPolicyCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_getPlaybackRestrictionPolicyCmd)
}
