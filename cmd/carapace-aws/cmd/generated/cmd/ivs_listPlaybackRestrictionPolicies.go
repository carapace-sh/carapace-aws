package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listPlaybackRestrictionPoliciesCmd = &cobra.Command{
	Use:   "list-playback-restriction-policies",
	Short: "Gets summary information about playback restriction policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listPlaybackRestrictionPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_listPlaybackRestrictionPoliciesCmd).Standalone()

		ivs_listPlaybackRestrictionPoliciesCmd.Flags().String("max-results", "", "Maximum number of policies to return.")
		ivs_listPlaybackRestrictionPoliciesCmd.Flags().String("next-token", "", "The first policy to retrieve.")
	})
	ivsCmd.AddCommand(ivs_listPlaybackRestrictionPoliciesCmd)
}
