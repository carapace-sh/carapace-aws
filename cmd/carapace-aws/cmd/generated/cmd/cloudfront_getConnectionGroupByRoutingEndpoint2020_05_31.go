package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getConnectionGroupByRoutingEndpoint2020_05_31Cmd = &cobra.Command{
	Use:   "get-connection-group-by-routing-endpoint2020_05_31",
	Short: "Gets information about a connection group by using the endpoint that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getConnectionGroupByRoutingEndpoint2020_05_31Cmd).Standalone()

	cloudfront_getConnectionGroupByRoutingEndpoint2020_05_31Cmd.Flags().String("routing-endpoint", "", "The routing endpoint for the target connection group, such as d111111abcdef8.cloudfront.net.")
	cloudfront_getConnectionGroupByRoutingEndpoint2020_05_31Cmd.MarkFlagRequired("routing-endpoint")
	cloudfrontCmd.AddCommand(cloudfront_getConnectionGroupByRoutingEndpoint2020_05_31Cmd)
}
