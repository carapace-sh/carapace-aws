package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_putPlaybackConfigurationCmd = &cobra.Command{
	Use:   "put-playback-configuration",
	Short: "Creates a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_putPlaybackConfigurationCmd).Standalone()

	mediatailor_putPlaybackConfigurationCmd.Flags().String("ad-conditioning-configuration", "", "The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns, and what priority MediaTailor uses when inserting ads.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("ad-decision-server-url", "", "The URL for the ad decision server (ADS).")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("avail-suppression", "", "The configuration for avail suppression, also known as ad suppression.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("bumper", "", "The configuration for bumpers.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("cdn-configuration", "", "The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("configuration-aliases", "", "The player parameters and aliases used as dynamic variables during session initialization.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("dash-configuration", "", "The configuration for DASH content.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("insertion-mode", "", "The setting that controls whether players can use stitched or guided ad insertion.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("live-pre-roll-configuration", "", "The configuration for pre-roll ad insertion.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("manifest-processing-rules", "", "The configuration for manifest processing rules.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("name", "", "The identifier for the playback configuration.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("personalization-threshold-seconds", "", "Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("slate-ad-url", "", "The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("tags", "", "The tags to assign to the playback configuration.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("transcode-profile-name", "", "The name that is used to associate this playback configuration with a custom transcode profile.")
	mediatailor_putPlaybackConfigurationCmd.Flags().String("video-content-source-url", "", "The URL prefix for the parent manifest for the stream, minus the asset ID.")
	mediatailor_putPlaybackConfigurationCmd.MarkFlagRequired("name")
	mediatailorCmd.AddCommand(mediatailor_putPlaybackConfigurationCmd)
}
