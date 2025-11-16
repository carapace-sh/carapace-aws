package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_updatePlaybackRestrictionPolicyCmd = &cobra.Command{
	Use:   "update-playback-restriction-policy",
	Short: "Updates a specified playback restriction policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_updatePlaybackRestrictionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_updatePlaybackRestrictionPolicyCmd).Standalone()

		ivs_updatePlaybackRestrictionPolicyCmd.Flags().String("allowed-countries", "", "A list of country codes that control geoblocking restriction.")
		ivs_updatePlaybackRestrictionPolicyCmd.Flags().String("allowed-origins", "", "A list of origin sites that control CORS restriction.")
		ivs_updatePlaybackRestrictionPolicyCmd.Flags().String("arn", "", "ARN of the playback-restriction-policy to be updated.")
		ivs_updatePlaybackRestrictionPolicyCmd.Flags().String("enable-strict-origin-enforcement", "", "Whether channel playback is constrained by origin site.")
		ivs_updatePlaybackRestrictionPolicyCmd.Flags().String("name", "", "Playback-restriction-policy name.")
		ivs_updatePlaybackRestrictionPolicyCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_updatePlaybackRestrictionPolicyCmd)
}
