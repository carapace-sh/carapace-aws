package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getCapacityManagerAttributesCmd = &cobra.Command{
	Use:   "get-capacity-manager-attributes",
	Short: "Retrieves the current configuration and status of EC2 Capacity Manager for your account, including enablement status, Organizations access settings, and data ingestion status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getCapacityManagerAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getCapacityManagerAttributesCmd).Standalone()

		ec2_getCapacityManagerAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerAttributesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getCapacityManagerAttributesCmd)
}
