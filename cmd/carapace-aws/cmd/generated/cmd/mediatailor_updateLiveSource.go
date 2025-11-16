package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_updateLiveSourceCmd = &cobra.Command{
	Use:   "update-live-source",
	Short: "Updates a live source's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_updateLiveSourceCmd).Standalone()

	mediatailor_updateLiveSourceCmd.Flags().String("http-package-configurations", "", "A list of HTTP package configurations for the live source on this account.")
	mediatailor_updateLiveSourceCmd.Flags().String("live-source-name", "", "The name of the live source.")
	mediatailor_updateLiveSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this Live Source.")
	mediatailor_updateLiveSourceCmd.MarkFlagRequired("http-package-configurations")
	mediatailor_updateLiveSourceCmd.MarkFlagRequired("live-source-name")
	mediatailor_updateLiveSourceCmd.MarkFlagRequired("source-location-name")
	mediatailorCmd.AddCommand(mediatailor_updateLiveSourceCmd)
}
