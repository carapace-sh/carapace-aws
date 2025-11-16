package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeDataSharesForConsumerCmd = &cobra.Command{
	Use:   "describe-data-shares-for-consumer",
	Short: "Returns a list of datashares where the account identifier being called is a consumer account identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeDataSharesForConsumerCmd).Standalone()

	redshift_describeDataSharesForConsumerCmd.Flags().String("consumer-arn", "", "The Amazon Resource Name (ARN) of the consumer namespace that returns in the list of datashares.")
	redshift_describeDataSharesForConsumerCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeDataSharesForConsumerCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeDataSharesForConsumerCmd.Flags().String("status", "", "An identifier giving the status of a datashare in the consumer cluster.")
	redshiftCmd.AddCommand(redshift_describeDataSharesForConsumerCmd)
}
