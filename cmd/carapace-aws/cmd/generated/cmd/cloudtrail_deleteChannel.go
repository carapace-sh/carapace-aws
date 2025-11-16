package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Deletes a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deleteChannelCmd).Standalone()

	cloudtrail_deleteChannelCmd.Flags().String("channel", "", "The ARN or the `UUID` value of the channel that you want to delete.")
	cloudtrail_deleteChannelCmd.MarkFlagRequired("channel")
	cloudtrailCmd.AddCommand(cloudtrail_deleteChannelCmd)
}
