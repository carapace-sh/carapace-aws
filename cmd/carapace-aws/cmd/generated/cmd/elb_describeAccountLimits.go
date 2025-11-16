package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeAccountLimitsCmd = &cobra.Command{
	Use:   "describe-account-limits",
	Short: "Describes the current Elastic Load Balancing resource limits for your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeAccountLimitsCmd).Standalone()

	elb_describeAccountLimitsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elb_describeAccountLimitsCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbCmd.AddCommand(elb_describeAccountLimitsCmd)
}
