package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Gets summary information about all channels in your account, in the Amazon Web Services region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_listChannelsCmd).Standalone()

		ivs_listChannelsCmd.Flags().String("filter-by-name", "", "Filters the channel list to match the specified name.")
		ivs_listChannelsCmd.Flags().String("filter-by-playback-restriction-policy-arn", "", "Filters the channel list to match the specified policy.")
		ivs_listChannelsCmd.Flags().String("filter-by-recording-configuration-arn", "", "Filters the channel list to match the specified recording-configuration ARN.")
		ivs_listChannelsCmd.Flags().String("max-results", "", "Maximum number of channels to return.")
		ivs_listChannelsCmd.Flags().String("next-token", "", "The first channel to retrieve.")
	})
	ivsCmd.AddCommand(ivs_listChannelsCmd)
}
