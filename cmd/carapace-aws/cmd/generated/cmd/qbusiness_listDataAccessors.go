package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listDataAccessorsCmd = &cobra.Command{
	Use:   "list-data-accessors",
	Short: "Lists the data accessors for a Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listDataAccessorsCmd).Standalone()

	qbusiness_listDataAccessorsCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
	qbusiness_listDataAccessorsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	qbusiness_listDataAccessorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qbusiness_listDataAccessorsCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_listDataAccessorsCmd)
}
