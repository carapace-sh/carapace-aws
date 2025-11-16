package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listTracksCmd = &cobra.Command{
	Use:   "list-tracks",
	Short: "List the Amazon Redshift Serverless versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listTracksCmd).Standalone()

	redshiftServerless_listTracksCmd.Flags().String("max-results", "", "The maximum number of response records to return in each call.")
	redshiftServerless_listTracksCmd.Flags().String("next-token", "", "If your initial `ListTracksRequest` operation returns a `nextToken`, you can include the returned `nextToken` in following `ListTracksRequest` operations, which returns results in the next page.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listTracksCmd)
}
