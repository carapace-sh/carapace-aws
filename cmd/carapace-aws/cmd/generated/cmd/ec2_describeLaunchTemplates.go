package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLaunchTemplatesCmd = &cobra.Command{
	Use:   "describe-launch-templates",
	Short: "Describes one or more launch templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLaunchTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeLaunchTemplatesCmd).Standalone()

		ec2_describeLaunchTemplatesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLaunchTemplatesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeLaunchTemplatesCmd.Flags().String("launch-template-ids", "", "One or more launch template IDs.")
		ec2_describeLaunchTemplatesCmd.Flags().String("launch-template-names", "", "One or more launch template names.")
		ec2_describeLaunchTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describeLaunchTemplatesCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		ec2_describeLaunchTemplatesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLaunchTemplatesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeLaunchTemplatesCmd)
}
