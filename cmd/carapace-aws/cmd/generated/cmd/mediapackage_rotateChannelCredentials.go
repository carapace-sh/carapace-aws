package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_rotateChannelCredentialsCmd = &cobra.Command{
	Use:   "rotate-channel-credentials",
	Short: "Changes the Channel's first IngestEndpoint's username and password.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_rotateChannelCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_rotateChannelCredentialsCmd).Standalone()

		mediapackage_rotateChannelCredentialsCmd.Flags().String("id", "", "The ID of the channel to update.")
		mediapackage_rotateChannelCredentialsCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_rotateChannelCredentialsCmd)
}
