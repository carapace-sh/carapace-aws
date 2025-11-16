package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBucketsCmd = &cobra.Command{
	Use:   "get-buckets",
	Short: "Returns information about one or more Amazon Lightsail buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBucketsCmd).Standalone()

	lightsail_getBucketsCmd.Flags().String("bucket-name", "", "The name of the bucket for which to return information.")
	lightsail_getBucketsCmd.Flags().String("include-connected-resources", "", "A Boolean value that indicates whether to include Lightsail instances that were given access to the bucket using the [SetResourceAccessForBucket](https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_SetResourceAccessForBucket.html) action.")
	lightsail_getBucketsCmd.Flags().String("include-cors", "", "A Boolean value that indicates whether to include Lightsail bucket CORS configuration in the response.")
	lightsail_getBucketsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getBucketsCmd)
}
