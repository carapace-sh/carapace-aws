package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeAccountLimitsCmd = &cobra.Command{
	Use:   "describe-account-limits",
	Short: "Describes the current Elastic Load Balancing resource limits for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeAccountLimitsCmd).Standalone()

	elbv2_describeAccountLimitsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elbv2_describeAccountLimitsCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbv2Cmd.AddCommand(elbv2_describeAccountLimitsCmd)
}
