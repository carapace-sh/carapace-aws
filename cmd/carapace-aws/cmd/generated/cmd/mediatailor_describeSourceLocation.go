package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_describeSourceLocationCmd = &cobra.Command{
	Use:   "describe-source-location",
	Short: "Describes a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_describeSourceLocationCmd).Standalone()

	mediatailor_describeSourceLocationCmd.Flags().String("source-location-name", "", "The name of the source location.")
	mediatailor_describeSourceLocationCmd.MarkFlagRequired("source-location-name")
	mediatailorCmd.AddCommand(mediatailor_describeSourceLocationCmd)
}
