package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyIpPoolsCmd = &cobra.Command{
	Use:   "modify-ip-pools",
	Short: "\\[Application Load Balancers] Modify the IP pool associated to a load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyIpPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_modifyIpPoolsCmd).Standalone()

		elbv2_modifyIpPoolsCmd.Flags().String("ipam-pools", "", "The IPAM pools to be modified.")
		elbv2_modifyIpPoolsCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
		elbv2_modifyIpPoolsCmd.Flags().String("remove-ipam-pools", "", "Remove the IP pools in use by the load balancer.")
		elbv2_modifyIpPoolsCmd.MarkFlagRequired("load-balancer-arn")
	})
	elbv2Cmd.AddCommand(elbv2_modifyIpPoolsCmd)
}
