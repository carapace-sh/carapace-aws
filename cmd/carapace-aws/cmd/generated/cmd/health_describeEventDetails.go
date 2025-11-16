package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEventDetailsCmd = &cobra.Command{
	Use:   "describe-event-details",
	Short: "Returns detailed information about one or more specified events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEventDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeEventDetailsCmd).Standalone()

		health_describeEventDetailsCmd.Flags().String("event-arns", "", "A list of event ARNs (unique identifiers).")
		health_describeEventDetailsCmd.Flags().String("locale", "", "The locale (language) to return information in.")
		health_describeEventDetailsCmd.MarkFlagRequired("event-arns")
	})
	healthCmd.AddCommand(health_describeEventDetailsCmd)
}
