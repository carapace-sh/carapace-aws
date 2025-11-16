package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyFleetCmd = &cobra.Command{
	Use:   "modify-fleet",
	Short: "Modifies the specified EC2 Fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyFleetCmd).Standalone()

	ec2_modifyFleetCmd.Flags().String("context", "", "Reserved.")
	ec2_modifyFleetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyFleetCmd.Flags().String("excess-capacity-termination-policy", "", "Indicates whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.")
	ec2_modifyFleetCmd.Flags().String("fleet-id", "", "The ID of the EC2 Fleet.")
	ec2_modifyFleetCmd.Flags().String("launch-template-configs", "", "The launch template and overrides.")
	ec2_modifyFleetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyFleetCmd.Flags().String("target-capacity-specification", "", "The size of the EC2 Fleet.")
	ec2_modifyFleetCmd.MarkFlagRequired("fleet-id")
	ec2_modifyFleetCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyFleetCmd)
}
