package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_cancelQueryCmd = &cobra.Command{
	Use:   "cancel-query",
	Short: "Cancels a query if the query is not in a terminated state, such as `CANCELLED`, `FAILED`, `TIMED_OUT`, or `FINISHED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_cancelQueryCmd).Standalone()

	cloudtrail_cancelQueryCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of an event data store on which the specified query is running.")
	cloudtrail_cancelQueryCmd.Flags().String("event-data-store-owner-account-id", "", "The account ID of the event data store owner.")
	cloudtrail_cancelQueryCmd.Flags().String("query-id", "", "The ID of the query that you want to cancel.")
	cloudtrail_cancelQueryCmd.MarkFlagRequired("query-id")
	cloudtrailCmd.AddCommand(cloudtrail_cancelQueryCmd)
}
