package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_getChannelPolicyCmd = &cobra.Command{
	Use:   "get-channel-policy",
	Short: "Returns the channel's IAM policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_getChannelPolicyCmd).Standalone()

	mediatailor_getChannelPolicyCmd.Flags().String("channel-name", "", "The name of the channel associated with this Channel Policy.")
	mediatailor_getChannelPolicyCmd.MarkFlagRequired("channel-name")
	mediatailorCmd.AddCommand(mediatailor_getChannelPolicyCmd)
}
