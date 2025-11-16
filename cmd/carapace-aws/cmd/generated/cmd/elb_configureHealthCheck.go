package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_configureHealthCheckCmd = &cobra.Command{
	Use:   "configure-health-check",
	Short: "Specifies the health check settings to use when evaluating the health state of your EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_configureHealthCheckCmd).Standalone()

	elb_configureHealthCheckCmd.Flags().String("health-check", "", "The configuration information.")
	elb_configureHealthCheckCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_configureHealthCheckCmd.MarkFlagRequired("health-check")
	elb_configureHealthCheckCmd.MarkFlagRequired("load-balancer-name")
	elbCmd.AddCommand(elb_configureHealthCheckCmd)
}
