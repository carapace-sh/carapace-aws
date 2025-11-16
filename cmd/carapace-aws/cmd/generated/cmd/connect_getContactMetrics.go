package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getContactMetricsCmd = &cobra.Command{
	Use:   "get-contact-metrics",
	Short: "Retrieves the position of the contact in the queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getContactMetricsCmd).Standalone()

	connect_getContactMetricsCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
	connect_getContactMetricsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_getContactMetricsCmd.Flags().String("metrics", "", "A list of contact-level metrics to retrieve.")
	connect_getContactMetricsCmd.MarkFlagRequired("contact-id")
	connect_getContactMetricsCmd.MarkFlagRequired("instance-id")
	connect_getContactMetricsCmd.MarkFlagRequired("metrics")
	connectCmd.AddCommand(connect_getContactMetricsCmd)
}
