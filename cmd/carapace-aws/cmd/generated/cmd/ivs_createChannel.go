package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a new channel and an associated stream key to start streaming.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_createChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_createChannelCmd).Standalone()

		ivs_createChannelCmd.Flags().Bool("authorized", false, "Whether the channel is private (enabled for playback authorization).")
		ivs_createChannelCmd.Flags().String("container-format", "", "Indicates which content-packaging format is used (MPEG-TS or fMP4).")
		ivs_createChannelCmd.Flags().Bool("insecure-ingest", false, "Whether the channel allows insecure RTMP and SRT ingest.")
		ivs_createChannelCmd.Flags().String("latency-mode", "", "Channel latency mode.")
		ivs_createChannelCmd.Flags().String("multitrack-input-configuration", "", "Object specifying multitrack input configuration.")
		ivs_createChannelCmd.Flags().String("name", "", "Channel name.")
		ivs_createChannelCmd.Flags().Bool("no-authorized", false, "Whether the channel is private (enabled for playback authorization).")
		ivs_createChannelCmd.Flags().Bool("no-insecure-ingest", false, "Whether the channel allows insecure RTMP and SRT ingest.")
		ivs_createChannelCmd.Flags().String("playback-restriction-policy-arn", "", "Playback-restriction-policy ARN.")
		ivs_createChannelCmd.Flags().String("preset", "", "Optional transcode preset for the channel.")
		ivs_createChannelCmd.Flags().String("recording-configuration-arn", "", "Recording-configuration ARN.")
		ivs_createChannelCmd.Flags().String("tags", "", "Array of 1-50 maps, each of the form `string:string (key:value)`.")
		ivs_createChannelCmd.Flags().String("type", "", "Channel type, which determines the allowable resolution and bitrate.")
		ivs_createChannelCmd.Flag("no-authorized").Hidden = true
		ivs_createChannelCmd.Flag("no-insecure-ingest").Hidden = true
	})
	ivsCmd.AddCommand(ivs_createChannelCmd)
}
