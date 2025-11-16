package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyHostsCmd = &cobra.Command{
	Use:   "modify-hosts",
	Short: "Modify the auto-placement setting of a Dedicated Host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyHostsCmd).Standalone()

		ec2_modifyHostsCmd.Flags().String("auto-placement", "", "Specify whether to enable or disable auto-placement.")
		ec2_modifyHostsCmd.Flags().String("host-ids", "", "The IDs of the Dedicated Hosts to modify.")
		ec2_modifyHostsCmd.Flags().String("host-maintenance", "", "Indicates whether to enable or disable host maintenance for the Dedicated Host.")
		ec2_modifyHostsCmd.Flags().String("host-recovery", "", "Indicates whether to enable or disable host recovery for the Dedicated Host.")
		ec2_modifyHostsCmd.Flags().String("instance-family", "", "Specifies the instance family to be supported by the Dedicated Host.")
		ec2_modifyHostsCmd.Flags().String("instance-type", "", "Specifies the instance type to be supported by the Dedicated Host.")
		ec2_modifyHostsCmd.MarkFlagRequired("host-ids")
	})
	ec2Cmd.AddCommand(ec2_modifyHostsCmd)
}
