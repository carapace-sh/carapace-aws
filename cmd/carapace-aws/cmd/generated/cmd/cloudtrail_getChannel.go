package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getChannelCmd = &cobra.Command{
	Use:   "get-channel",
	Short: "Returns information about a specific channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getChannelCmd).Standalone()

	cloudtrail_getChannelCmd.Flags().String("channel", "", "The ARN or `UUID` of a channel.")
	cloudtrail_getChannelCmd.MarkFlagRequired("channel")
	cloudtrailCmd.AddCommand(cloudtrail_getChannelCmd)
}
