package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeDataSharesForProducerCmd = &cobra.Command{
	Use:   "describe-data-shares-for-producer",
	Short: "Returns a list of datashares when the account identifier being called is a producer account identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeDataSharesForProducerCmd).Standalone()

	redshift_describeDataSharesForProducerCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeDataSharesForProducerCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeDataSharesForProducerCmd.Flags().String("producer-arn", "", "The Amazon Resource Name (ARN) of the producer namespace that returns in the list of datashares.")
	redshift_describeDataSharesForProducerCmd.Flags().String("status", "", "An identifier giving the status of a datashare in the producer.")
	redshiftCmd.AddCommand(redshift_describeDataSharesForProducerCmd)
}
