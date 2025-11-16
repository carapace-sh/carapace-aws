package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createLocationCmd = &cobra.Command{
	Use:   "create-location",
	Short: "**This API works with the following fleet types:** Anywhere\n\nCreates a custom location for use in an Anywhere fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createLocationCmd).Standalone()

		gamelift_createLocationCmd.Flags().String("location-name", "", "A descriptive name for the custom location.")
		gamelift_createLocationCmd.Flags().String("tags", "", "A list of labels to assign to the new resource.")
		gamelift_createLocationCmd.MarkFlagRequired("location-name")
	})
	gameliftCmd.AddCommand(gamelift_createLocationCmd)
}
