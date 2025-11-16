package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteCoipPoolCmd = &cobra.Command{
	Use:   "delete-coip-pool",
	Short: "Deletes a pool of customer-owned IP (CoIP) addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteCoipPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteCoipPoolCmd).Standalone()

		ec2_deleteCoipPoolCmd.Flags().String("coip-pool-id", "", "The ID of the CoIP pool that you want to delete.")
		ec2_deleteCoipPoolCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCoipPoolCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCoipPoolCmd.MarkFlagRequired("coip-pool-id")
		ec2_deleteCoipPoolCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteCoipPoolCmd)
}
