package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_allocateHostsCmd = &cobra.Command{
	Use:   "allocate-hosts",
	Short: "Allocates a Dedicated Host to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_allocateHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_allocateHostsCmd).Standalone()

		ec2_allocateHostsCmd.Flags().String("asset-ids", "", "The IDs of the Outpost hardware assets on which to allocate the Dedicated Hosts.")
		ec2_allocateHostsCmd.Flags().String("auto-placement", "", "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.")
		ec2_allocateHostsCmd.Flags().String("availability-zone", "", "The Availability Zone in which to allocate the Dedicated Host.")
		ec2_allocateHostsCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone.")
		ec2_allocateHostsCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_allocateHostsCmd.Flags().String("host-maintenance", "", "Indicates whether to enable or disable host maintenance for the Dedicated Host.")
		ec2_allocateHostsCmd.Flags().String("host-recovery", "", "Indicates whether to enable or disable host recovery for the Dedicated Host.")
		ec2_allocateHostsCmd.Flags().String("instance-family", "", "Specifies the instance family to be supported by the Dedicated Hosts.")
		ec2_allocateHostsCmd.Flags().String("instance-type", "", "Specifies the instance type to be supported by the Dedicated Hosts.")
		ec2_allocateHostsCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.")
		ec2_allocateHostsCmd.Flags().String("quantity", "", "The number of Dedicated Hosts to allocate to your account with these parameters.")
		ec2_allocateHostsCmd.Flags().String("tag-specifications", "", "The tags to apply to the Dedicated Host during creation.")
	})
	ec2Cmd.AddCommand(ec2_allocateHostsCmd)
}
