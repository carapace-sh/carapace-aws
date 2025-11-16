package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTrafficMirrorFilterRuleCmd = &cobra.Command{
	Use:   "modify-traffic-mirror-filter-rule",
	Short: "Modifies the specified Traffic Mirror rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTrafficMirrorFilterRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTrafficMirrorFilterRuleCmd).Standalone()

		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("description", "", "The description to assign to the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("destination-cidr-block", "", "The destination CIDR block to assign to the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("destination-port-range", "", "The destination ports that are associated with the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("protocol", "", "The protocol, for example TCP, to assign to the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("remove-fields", "", "The properties that you want to remove from the Traffic Mirror filter rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("rule-action", "", "The action to assign to the rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("rule-number", "", "The number of the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("source-cidr-block", "", "The source CIDR block to assign to the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("source-port-range", "", "The port range to assign to the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("traffic-direction", "", "The type of traffic to assign to the rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flags().String("traffic-mirror-filter-rule-id", "", "The ID of the Traffic Mirror rule.")
		ec2_modifyTrafficMirrorFilterRuleCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTrafficMirrorFilterRuleCmd.MarkFlagRequired("traffic-mirror-filter-rule-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTrafficMirrorFilterRuleCmd)
}
