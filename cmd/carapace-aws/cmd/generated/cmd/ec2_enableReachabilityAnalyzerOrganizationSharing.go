package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableReachabilityAnalyzerOrganizationSharingCmd = &cobra.Command{
	Use:   "enable-reachability-analyzer-organization-sharing",
	Short: "Establishes a trust relationship between Reachability Analyzer and Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableReachabilityAnalyzerOrganizationSharingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableReachabilityAnalyzerOrganizationSharingCmd).Standalone()

		ec2_enableReachabilityAnalyzerOrganizationSharingCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableReachabilityAnalyzerOrganizationSharingCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableReachabilityAnalyzerOrganizationSharingCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableReachabilityAnalyzerOrganizationSharingCmd)
}
