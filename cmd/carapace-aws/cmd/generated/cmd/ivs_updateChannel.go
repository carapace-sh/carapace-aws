package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Updates a channel's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_updateChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_updateChannelCmd).Standalone()

		ivs_updateChannelCmd.Flags().String("arn", "", "ARN of the channel to be updated.")
		ivs_updateChannelCmd.Flags().Bool("authorized", false, "Whether the channel is private (enabled for playback authorization).")
		ivs_updateChannelCmd.Flags().String("container-format", "", "Indicates which content-packaging format is used (MPEG-TS or fMP4).")
		ivs_updateChannelCmd.Flags().Bool("insecure-ingest", false, "Whether the channel allows insecure RTMP and SRT ingest.")
		ivs_updateChannelCmd.Flags().String("latency-mode", "", "Channel latency mode.")
		ivs_updateChannelCmd.Flags().String("multitrack-input-configuration", "", "Object specifying multitrack input configuration.")
		ivs_updateChannelCmd.Flags().String("name", "", "Channel name.")
		ivs_updateChannelCmd.Flags().Bool("no-authorized", false, "Whether the channel is private (enabled for playback authorization).")
		ivs_updateChannelCmd.Flags().Bool("no-insecure-ingest", false, "Whether the channel allows insecure RTMP and SRT ingest.")
		ivs_updateChannelCmd.Flags().String("playback-restriction-policy-arn", "", "Playback-restriction-policy ARN.")
		ivs_updateChannelCmd.Flags().String("preset", "", "Optional transcode preset for the channel.")
		ivs_updateChannelCmd.Flags().String("recording-configuration-arn", "", "Recording-configuration ARN.")
		ivs_updateChannelCmd.Flags().String("type", "", "Channel type, which determines the allowable resolution and bitrate.")
		ivs_updateChannelCmd.MarkFlagRequired("arn")
		ivs_updateChannelCmd.Flag("no-authorized").Hidden = true
		ivs_updateChannelCmd.Flag("no-insecure-ingest").Hidden = true
	})
	ivsCmd.AddCommand(ivs_updateChannelCmd)
}
