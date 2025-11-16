package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outposts_listSharedEndpointsCmd = &cobra.Command{
	Use:   "list-shared-endpoints",
	Short: "Lists all endpoints associated with an Outpost that has been shared by Amazon Web Services Resource Access Manager (RAM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outposts_listSharedEndpointsCmd).Standalone()

	s3outposts_listSharedEndpointsCmd.Flags().String("max-results", "", "The maximum number of endpoints that will be returned in the response.")
	s3outposts_listSharedEndpointsCmd.Flags().String("next-token", "", "If a previous response from this operation included a `NextToken` value, you can provide that value here to retrieve the next page of results.")
	s3outposts_listSharedEndpointsCmd.Flags().String("outpost-id", "", "The ID of the Amazon Web Services Outpost.")
	s3outposts_listSharedEndpointsCmd.MarkFlagRequired("outpost-id")
	s3outpostsCmd.AddCommand(s3outposts_listSharedEndpointsCmd)
}
