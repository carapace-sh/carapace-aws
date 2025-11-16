package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeDataSharesCmd = &cobra.Command{
	Use:   "describe-data-shares",
	Short: "Shows the status of any inbound or outbound datashares available in the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeDataSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeDataSharesCmd).Standalone()

		redshift_describeDataSharesCmd.Flags().String("data-share-arn", "", "The Amazon resource name (ARN) of the datashare to describe details of.")
		redshift_describeDataSharesCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeDataSharesCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	})
	redshiftCmd.AddCommand(redshift_describeDataSharesCmd)
}
