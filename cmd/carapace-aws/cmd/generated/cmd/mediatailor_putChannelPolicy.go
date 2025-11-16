package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_putChannelPolicyCmd = &cobra.Command{
	Use:   "put-channel-policy",
	Short: "Creates an IAM policy for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_putChannelPolicyCmd).Standalone()

	mediatailor_putChannelPolicyCmd.Flags().String("channel-name", "", "The channel name associated with this Channel Policy.")
	mediatailor_putChannelPolicyCmd.Flags().String("policy", "", "Adds an IAM role that determines the permissions of your channel.")
	mediatailor_putChannelPolicyCmd.MarkFlagRequired("channel-name")
	mediatailor_putChannelPolicyCmd.MarkFlagRequired("policy")
	mediatailorCmd.AddCommand(mediatailor_putChannelPolicyCmd)
}
