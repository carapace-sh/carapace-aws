package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTrafficMirrorFilterRuleCmd = &cobra.Command{
	Use:   "create-traffic-mirror-filter-rule",
	Short: "Creates a Traffic Mirror filter rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTrafficMirrorFilterRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTrafficMirrorFilterRuleCmd).Standalone()

		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("description", "", "The description of the Traffic Mirror rule.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("destination-cidr-block", "", "The destination CIDR block to assign to the Traffic Mirror rule.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("destination-port-range", "", "The destination port range.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("protocol", "", "The protocol, for example UDP, to assign to the Traffic Mirror rule.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("rule-action", "", "The action to take on the filtered traffic.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("rule-number", "", "The number of the Traffic Mirror rule.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("source-cidr-block", "", "The source CIDR block to assign to the Traffic Mirror rule.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("source-port-range", "", "The source port range.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("tag-specifications", "", "Traffic Mirroring tags specifications.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("traffic-direction", "", "The type of traffic.")
		ec2_createTrafficMirrorFilterRuleCmd.Flags().String("traffic-mirror-filter-id", "", "The ID of the filter that this rule is associated with.")
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("destination-cidr-block")
		ec2_createTrafficMirrorFilterRuleCmd.Flag("no-dry-run").Hidden = true
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("rule-action")
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("rule-number")
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("source-cidr-block")
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("traffic-direction")
		ec2_createTrafficMirrorFilterRuleCmd.MarkFlagRequired("traffic-mirror-filter-id")
	})
	ec2Cmd.AddCommand(ec2_createTrafficMirrorFilterRuleCmd)
}
