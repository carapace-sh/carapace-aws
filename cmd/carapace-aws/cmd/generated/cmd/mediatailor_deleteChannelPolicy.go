package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteChannelPolicyCmd = &cobra.Command{
	Use:   "delete-channel-policy",
	Short: "The channel policy to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteChannelPolicyCmd).Standalone()

	mediatailor_deleteChannelPolicyCmd.Flags().String("channel-name", "", "The name of the channel associated with this channel policy.")
	mediatailor_deleteChannelPolicyCmd.MarkFlagRequired("channel-name")
	mediatailorCmd.AddCommand(mediatailor_deleteChannelPolicyCmd)
}
