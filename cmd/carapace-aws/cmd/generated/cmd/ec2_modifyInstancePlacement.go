package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstancePlacementCmd = &cobra.Command{
	Use:   "modify-instance-placement",
	Short: "Modifies the placement attributes for a specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstancePlacementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyInstancePlacementCmd).Standalone()

		ec2_modifyInstancePlacementCmd.Flags().String("affinity", "", "The affinity setting for the instance.")
		ec2_modifyInstancePlacementCmd.Flags().String("group-id", "", "The Group Id of a placement group.")
		ec2_modifyInstancePlacementCmd.Flags().String("group-name", "", "The name of the placement group in which to place the instance.")
		ec2_modifyInstancePlacementCmd.Flags().String("host-id", "", "The ID of the Dedicated Host with which to associate the instance.")
		ec2_modifyInstancePlacementCmd.Flags().String("host-resource-group-arn", "", "The ARN of the host resource group in which to place the instance.")
		ec2_modifyInstancePlacementCmd.Flags().String("instance-id", "", "The ID of the instance that you are modifying.")
		ec2_modifyInstancePlacementCmd.Flags().String("partition-number", "", "The number of the partition in which to place the instance.")
		ec2_modifyInstancePlacementCmd.Flags().String("tenancy", "", "The tenancy for the instance.")
		ec2_modifyInstancePlacementCmd.MarkFlagRequired("instance-id")
	})
	ec2Cmd.AddCommand(ec2_modifyInstancePlacementCmd)
}
