package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteSourceLocationCmd = &cobra.Command{
	Use:   "delete-source-location",
	Short: "Deletes a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteSourceLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_deleteSourceLocationCmd).Standalone()

		mediatailor_deleteSourceLocationCmd.Flags().String("source-location-name", "", "The name of the source location.")
		mediatailor_deleteSourceLocationCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_deleteSourceLocationCmd)
}
