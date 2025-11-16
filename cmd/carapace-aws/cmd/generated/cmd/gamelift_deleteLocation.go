package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteLocationCmd = &cobra.Command{
	Use:   "delete-location",
	Short: "**This API works with the following fleet types:** Anywhere\n\nDeletes a custom location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteLocationCmd).Standalone()

		gamelift_deleteLocationCmd.Flags().String("location-name", "", "The location name of the custom location to be deleted.")
		gamelift_deleteLocationCmd.MarkFlagRequired("location-name")
	})
	gameliftCmd.AddCommand(gamelift_deleteLocationCmd)
}
