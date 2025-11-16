package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createLiveSourceCmd = &cobra.Command{
	Use:   "create-live-source",
	Short: "The live source configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createLiveSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_createLiveSourceCmd).Standalone()

		mediatailor_createLiveSourceCmd.Flags().String("http-package-configurations", "", "A list of HTTP package configuration parameters for this live source.")
		mediatailor_createLiveSourceCmd.Flags().String("live-source-name", "", "The name of the live source.")
		mediatailor_createLiveSourceCmd.Flags().String("source-location-name", "", "The name of the source location.")
		mediatailor_createLiveSourceCmd.Flags().String("tags", "", "The tags to assign to the live source.")
		mediatailor_createLiveSourceCmd.MarkFlagRequired("http-package-configurations")
		mediatailor_createLiveSourceCmd.MarkFlagRequired("live-source-name")
		mediatailor_createLiveSourceCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_createLiveSourceCmd)
}
