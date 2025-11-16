package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteLiveSourceCmd = &cobra.Command{
	Use:   "delete-live-source",
	Short: "The live source to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteLiveSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_deleteLiveSourceCmd).Standalone()

		mediatailor_deleteLiveSourceCmd.Flags().String("live-source-name", "", "The name of the live source.")
		mediatailor_deleteLiveSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this Live Source.")
		mediatailor_deleteLiveSourceCmd.MarkFlagRequired("live-source-name")
		mediatailor_deleteLiveSourceCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_deleteLiveSourceCmd)
}
