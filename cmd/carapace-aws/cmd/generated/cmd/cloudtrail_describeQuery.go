package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_describeQueryCmd = &cobra.Command{
	Use:   "describe-query",
	Short: "Returns metadata about a query, including query run time in milliseconds, number of events scanned and matched, and query status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_describeQueryCmd).Standalone()

	cloudtrail_describeQueryCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of an event data store on which the specified query was run.")
	cloudtrail_describeQueryCmd.Flags().String("event-data-store-owner-account-id", "", "The account ID of the event data store owner.")
	cloudtrail_describeQueryCmd.Flags().String("query-alias", "", "The alias that identifies a query template.")
	cloudtrail_describeQueryCmd.Flags().String("query-id", "", "The query ID.")
	cloudtrail_describeQueryCmd.Flags().String("refresh-id", "", "The ID of the dashboard refresh.")
	cloudtrailCmd.AddCommand(cloudtrail_describeQueryCmd)
}
