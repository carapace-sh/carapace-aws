package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTrafficMirrorFilterRulesCmd = &cobra.Command{
	Use:   "describe-traffic-mirror-filter-rules",
	Short: "Describe traffic mirror filters that determine the traffic that is mirrored.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTrafficMirrorFilterRulesCmd).Standalone()

	ec2_describeTrafficMirrorFilterRulesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().String("filters", "", "Traffic mirror filters.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().String("traffic-mirror-filter-id", "", "Traffic filter ID.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flags().String("traffic-mirror-filter-rule-ids", "", "Traffic filter rule IDs.")
	ec2_describeTrafficMirrorFilterRulesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeTrafficMirrorFilterRulesCmd)
}
