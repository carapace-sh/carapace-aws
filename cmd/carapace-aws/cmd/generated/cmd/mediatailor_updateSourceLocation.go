package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_updateSourceLocationCmd = &cobra.Command{
	Use:   "update-source-location",
	Short: "Updates a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_updateSourceLocationCmd).Standalone()

	mediatailor_updateSourceLocationCmd.Flags().String("access-configuration", "", "Access configuration parameters.")
	mediatailor_updateSourceLocationCmd.Flags().String("default-segment-delivery-configuration", "", "The optional configuration for the host server that serves segments.")
	mediatailor_updateSourceLocationCmd.Flags().String("http-configuration", "", "The HTTP configuration for the source location.")
	mediatailor_updateSourceLocationCmd.Flags().String("segment-delivery-configurations", "", "A list of the segment delivery configurations associated with this resource.")
	mediatailor_updateSourceLocationCmd.Flags().String("source-location-name", "", "The name of the source location.")
	mediatailor_updateSourceLocationCmd.MarkFlagRequired("http-configuration")
	mediatailor_updateSourceLocationCmd.MarkFlagRequired("source-location-name")
	mediatailorCmd.AddCommand(mediatailor_updateSourceLocationCmd)
}
