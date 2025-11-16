package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTrafficMirrorFilterRuleCmd = &cobra.Command{
	Use:   "delete-traffic-mirror-filter-rule",
	Short: "Deletes the specified Traffic Mirror rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTrafficMirrorFilterRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTrafficMirrorFilterRuleCmd).Standalone()

		ec2_deleteTrafficMirrorFilterRuleCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorFilterRuleCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorFilterRuleCmd.Flags().String("traffic-mirror-filter-rule-id", "", "The ID of the Traffic Mirror rule.")
		ec2_deleteTrafficMirrorFilterRuleCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTrafficMirrorFilterRuleCmd.MarkFlagRequired("traffic-mirror-filter-rule-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTrafficMirrorFilterRuleCmd)
}
