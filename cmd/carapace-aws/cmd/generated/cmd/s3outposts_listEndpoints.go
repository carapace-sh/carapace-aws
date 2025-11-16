package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outposts_listEndpointsCmd = &cobra.Command{
	Use:   "list-endpoints",
	Short: "Lists endpoints associated with the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outposts_listEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3outposts_listEndpointsCmd).Standalone()

		s3outposts_listEndpointsCmd.Flags().String("max-results", "", "The maximum number of endpoints that will be returned in the response.")
		s3outposts_listEndpointsCmd.Flags().String("next-token", "", "If a previous response from this operation included a `NextToken` value, provide that value here to retrieve the next page of results.")
	})
	s3outpostsCmd.AddCommand(s3outposts_listEndpointsCmd)
}
