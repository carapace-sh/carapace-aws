package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describePoliciesCmd = &cobra.Command{
	Use:   "describe-policies",
	Short: "Gets information about the scaling policies in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describePoliciesCmd).Standalone()

	autoscaling_describePoliciesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_describePoliciesCmd.Flags().String("max-records", "", "The maximum number of items to be returned with each call.")
	autoscaling_describePoliciesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	autoscaling_describePoliciesCmd.Flags().String("policy-names", "", "The names of one or more policies.")
	autoscaling_describePoliciesCmd.Flags().String("policy-types", "", "One or more policy types.")
	autoscalingCmd.AddCommand(autoscaling_describePoliciesCmd)
}
