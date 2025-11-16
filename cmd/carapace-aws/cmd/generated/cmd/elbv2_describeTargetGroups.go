package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTargetGroupsCmd = &cobra.Command{
	Use:   "describe-target-groups",
	Short: "Describes the specified target groups or all of your target groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTargetGroupsCmd).Standalone()

	elbv2_describeTargetGroupsCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_describeTargetGroupsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elbv2_describeTargetGroupsCmd.Flags().String("names", "", "The names of the target groups.")
	elbv2_describeTargetGroupsCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbv2_describeTargetGroupsCmd.Flags().String("target-group-arns", "", "The Amazon Resource Names (ARN) of the target groups.")
	elbv2Cmd.AddCommand(elbv2_describeTargetGroupsCmd)
}
