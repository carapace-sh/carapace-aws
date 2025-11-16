package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyReservedInstancesCmd = &cobra.Command{
	Use:   "modify-reserved-instances",
	Short: "Modifies the configuration of your Reserved Instances, such as the Availability Zone, instance count, or instance type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyReservedInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyReservedInstancesCmd).Standalone()

		ec2_modifyReservedInstancesCmd.Flags().String("client-token", "", "A unique, case-sensitive token you provide to ensure idempotency of your modification request.")
		ec2_modifyReservedInstancesCmd.Flags().String("reserved-instances-ids", "", "The IDs of the Reserved Instances to modify.")
		ec2_modifyReservedInstancesCmd.Flags().String("target-configurations", "", "The configuration settings for the Reserved Instances to modify.")
		ec2_modifyReservedInstancesCmd.MarkFlagRequired("reserved-instances-ids")
		ec2_modifyReservedInstancesCmd.MarkFlagRequired("target-configurations")
	})
	ec2Cmd.AddCommand(ec2_modifyReservedInstancesCmd)
}
