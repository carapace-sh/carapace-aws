package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCoipPoolCmd = &cobra.Command{
	Use:   "create-coip-pool",
	Short: "Creates a pool of customer-owned IP (CoIP) addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCoipPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createCoipPoolCmd).Standalone()

		ec2_createCoipPoolCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCoipPoolCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
		ec2_createCoipPoolCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCoipPoolCmd.Flags().String("tag-specifications", "", "The tags to assign to the CoIP address pool.")
		ec2_createCoipPoolCmd.MarkFlagRequired("local-gateway-route-table-id")
		ec2_createCoipPoolCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createCoipPoolCmd)
}
