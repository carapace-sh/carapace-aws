package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_describeLiveSourceCmd = &cobra.Command{
	Use:   "describe-live-source",
	Short: "The live source to describe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_describeLiveSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_describeLiveSourceCmd).Standalone()

		mediatailor_describeLiveSourceCmd.Flags().String("live-source-name", "", "The name of the live source.")
		mediatailor_describeLiveSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this Live Source.")
		mediatailor_describeLiveSourceCmd.MarkFlagRequired("live-source-name")
		mediatailor_describeLiveSourceCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_describeLiveSourceCmd)
}
