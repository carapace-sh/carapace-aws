package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_startQueryCmd = &cobra.Command{
	Use:   "start-query",
	Short: "Starts a CloudTrail Lake query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_startQueryCmd).Standalone()

	cloudtrail_startQueryCmd.Flags().String("delivery-s3-uri", "", "The URI for the S3 bucket where CloudTrail delivers the query results.")
	cloudtrail_startQueryCmd.Flags().String("event-data-store-owner-account-id", "", "The account ID of the event data store owner.")
	cloudtrail_startQueryCmd.Flags().String("query-alias", "", "The alias that identifies a query template.")
	cloudtrail_startQueryCmd.Flags().String("query-parameters", "", "The query parameters for the specified `QueryAlias`.")
	cloudtrail_startQueryCmd.Flags().String("query-statement", "", "The SQL code of your query.")
	cloudtrailCmd.AddCommand(cloudtrail_startQueryCmd)
}
