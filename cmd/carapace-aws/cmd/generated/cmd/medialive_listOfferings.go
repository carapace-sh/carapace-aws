package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listOfferingsCmd = &cobra.Command{
	Use:   "list-offerings",
	Short: "List offerings available for purchase.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listOfferingsCmd).Standalone()

	medialive_listOfferingsCmd.Flags().String("channel-class", "", "Filter by channel class, 'STANDARD' or 'SINGLE\\_PIPELINE'")
	medialive_listOfferingsCmd.Flags().String("channel-configuration", "", "Filter to offerings that match the configuration of an existing channel, e.g. '2345678' (a channel ID)")
	medialive_listOfferingsCmd.Flags().String("codec", "", "Filter by codec, 'AVC', 'HEVC', 'MPEG2', 'AUDIO', 'LINK', or 'AV1'")
	medialive_listOfferingsCmd.Flags().String("duration", "", "Filter by offering duration, e.g. '12'")
	medialive_listOfferingsCmd.Flags().String("max-results", "", "")
	medialive_listOfferingsCmd.Flags().String("maximum-bitrate", "", "Filter by bitrate, 'MAX\\_10\\_MBPS', 'MAX\\_20\\_MBPS', or 'MAX\\_50\\_MBPS'")
	medialive_listOfferingsCmd.Flags().String("maximum-framerate", "", "Filter by framerate, 'MAX\\_30\\_FPS' or 'MAX\\_60\\_FPS'")
	medialive_listOfferingsCmd.Flags().String("next-token", "", "")
	medialive_listOfferingsCmd.Flags().String("resolution", "", "Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'")
	medialive_listOfferingsCmd.Flags().String("resource-type", "", "Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'")
	medialive_listOfferingsCmd.Flags().String("special-feature", "", "Filter by special feature, 'ADVANCED\\_AUDIO' or 'AUDIO\\_NORMALIZATION'")
	medialive_listOfferingsCmd.Flags().String("video-quality", "", "Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'")
	medialiveCmd.AddCommand(medialive_listOfferingsCmd)
}
