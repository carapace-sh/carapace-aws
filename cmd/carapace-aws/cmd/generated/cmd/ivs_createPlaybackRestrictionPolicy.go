package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_createPlaybackRestrictionPolicyCmd = &cobra.Command{
	Use:   "create-playback-restriction-policy",
	Short: "Creates a new playback restriction policy, for constraining playback by countries and/or origins.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_createPlaybackRestrictionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_createPlaybackRestrictionPolicyCmd).Standalone()

		ivs_createPlaybackRestrictionPolicyCmd.Flags().String("allowed-countries", "", "A list of country codes that control geoblocking restriction.")
		ivs_createPlaybackRestrictionPolicyCmd.Flags().String("allowed-origins", "", "A list of origin sites that control CORS restriction.")
		ivs_createPlaybackRestrictionPolicyCmd.Flags().String("enable-strict-origin-enforcement", "", "Whether channel playback is constrained by origin site.")
		ivs_createPlaybackRestrictionPolicyCmd.Flags().String("name", "", "Playback-restriction-policy name.")
		ivs_createPlaybackRestrictionPolicyCmd.Flags().String("tags", "", "Array of 1-50 maps, each of the form `string:string (key:value)`.")
	})
	ivsCmd.AddCommand(ivs_createPlaybackRestrictionPolicyCmd)
}
