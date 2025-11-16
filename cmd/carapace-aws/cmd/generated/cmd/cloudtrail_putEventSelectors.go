package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_putEventSelectorsCmd = &cobra.Command{
	Use:   "put-event-selectors",
	Short: "Configures event selectors (also referred to as *basic event selectors*) or advanced event selectors for your trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_putEventSelectorsCmd).Standalone()

	cloudtrail_putEventSelectorsCmd.Flags().String("advanced-event-selectors", "", "Specifies the settings for advanced event selectors.")
	cloudtrail_putEventSelectorsCmd.Flags().String("event-selectors", "", "Specifies the settings for your event selectors.")
	cloudtrail_putEventSelectorsCmd.Flags().String("trail-name", "", "Specifies the name of the trail or trail ARN.")
	cloudtrail_putEventSelectorsCmd.MarkFlagRequired("trail-name")
	cloudtrailCmd.AddCommand(cloudtrail_putEventSelectorsCmd)
}
