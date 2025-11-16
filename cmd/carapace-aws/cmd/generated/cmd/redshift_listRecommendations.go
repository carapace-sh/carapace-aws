package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "List the Amazon Redshift Advisor recommendations for one or multiple Amazon Redshift clusters in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_listRecommendationsCmd).Standalone()

	redshift_listRecommendationsCmd.Flags().String("cluster-identifier", "", "The unique identifier of the Amazon Redshift cluster for which the list of Advisor recommendations is returned.")
	redshift_listRecommendationsCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
	redshift_listRecommendationsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_listRecommendationsCmd.Flags().String("namespace-arn", "", "The Amazon Redshift cluster namespace Amazon Resource Name (ARN) for which the list of Advisor recommendations is returned.")
	redshiftCmd.AddCommand(redshift_listRecommendationsCmd)
}
