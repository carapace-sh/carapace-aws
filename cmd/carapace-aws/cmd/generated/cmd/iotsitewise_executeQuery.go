package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_executeQueryCmd = &cobra.Command{
	Use:   "execute-query",
	Short: "Run SQL queries to retrieve metadata and time-series data from asset models, assets, measurements, metrics, transforms, and aggregates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_executeQueryCmd).Standalone()

	iotsitewise_executeQueryCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_executeQueryCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotsitewise_executeQueryCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotsitewise_executeQueryCmd.Flags().String("query-statement", "", "The IoT SiteWise query statement.")
	iotsitewise_executeQueryCmd.MarkFlagRequired("query-statement")
	iotsitewiseCmd.AddCommand(iotsitewise_executeQueryCmd)
}
