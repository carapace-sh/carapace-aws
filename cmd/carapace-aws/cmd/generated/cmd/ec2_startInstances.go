package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_startInstancesCmd = &cobra.Command{
	Use:   "start-instances",
	Short: "Starts an Amazon EBS-backed instance that you've previously stopped.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_startInstancesCmd).Standalone()

	ec2_startInstancesCmd.Flags().String("additional-info", "", "Reserved.")
	ec2_startInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_startInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
	ec2_startInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_startInstancesCmd.MarkFlagRequired("instance-ids")
	ec2_startInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_startInstancesCmd)
}
