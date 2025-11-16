package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createSourceLocationCmd = &cobra.Command{
	Use:   "create-source-location",
	Short: "Creates a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createSourceLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_createSourceLocationCmd).Standalone()

		mediatailor_createSourceLocationCmd.Flags().String("access-configuration", "", "Access configuration parameters.")
		mediatailor_createSourceLocationCmd.Flags().String("default-segment-delivery-configuration", "", "The optional configuration for the server that serves segments.")
		mediatailor_createSourceLocationCmd.Flags().String("http-configuration", "", "The source's HTTP package configurations.")
		mediatailor_createSourceLocationCmd.Flags().String("segment-delivery-configurations", "", "A list of the segment delivery configurations associated with this resource.")
		mediatailor_createSourceLocationCmd.Flags().String("source-location-name", "", "The name associated with the source location.")
		mediatailor_createSourceLocationCmd.Flags().String("tags", "", "The tags to assign to the source location.")
		mediatailor_createSourceLocationCmd.MarkFlagRequired("http-configuration")
		mediatailor_createSourceLocationCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_createSourceLocationCmd)
}
