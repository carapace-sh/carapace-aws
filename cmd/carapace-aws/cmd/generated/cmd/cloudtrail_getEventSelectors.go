package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getEventSelectorsCmd = &cobra.Command{
	Use:   "get-event-selectors",
	Short: "Describes the settings for the event selectors that you configured for your trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getEventSelectorsCmd).Standalone()

	cloudtrail_getEventSelectorsCmd.Flags().String("trail-name", "", "Specifies the name of the trail or trail ARN.")
	cloudtrail_getEventSelectorsCmd.MarkFlagRequired("trail-name")
	cloudtrailCmd.AddCommand(cloudtrail_getEventSelectorsCmd)
}
