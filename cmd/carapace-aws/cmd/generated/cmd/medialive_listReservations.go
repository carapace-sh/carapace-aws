package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listReservationsCmd = &cobra.Command{
	Use:   "list-reservations",
	Short: "List purchased reservations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listReservationsCmd).Standalone()

	medialive_listReservationsCmd.Flags().String("channel-class", "", "Filter by channel class, 'STANDARD' or 'SINGLE\\_PIPELINE'")
	medialive_listReservationsCmd.Flags().String("codec", "", "Filter by codec, 'AVC', 'HEVC', 'MPEG2', 'AUDIO', 'LINK', or 'AV1'")
	medialive_listReservationsCmd.Flags().String("max-results", "", "")
	medialive_listReservationsCmd.Flags().String("maximum-bitrate", "", "Filter by bitrate, 'MAX\\_10\\_MBPS', 'MAX\\_20\\_MBPS', or 'MAX\\_50\\_MBPS'")
	medialive_listReservationsCmd.Flags().String("maximum-framerate", "", "Filter by framerate, 'MAX\\_30\\_FPS' or 'MAX\\_60\\_FPS'")
	medialive_listReservationsCmd.Flags().String("next-token", "", "")
	medialive_listReservationsCmd.Flags().String("resolution", "", "Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'")
	medialive_listReservationsCmd.Flags().String("resource-type", "", "Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'")
	medialive_listReservationsCmd.Flags().String("special-feature", "", "Filter by special feature, 'ADVANCED\\_AUDIO' or 'AUDIO\\_NORMALIZATION'")
	medialive_listReservationsCmd.Flags().String("video-quality", "", "Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'")
	medialiveCmd.AddCommand(medialive_listReservationsCmd)
}
