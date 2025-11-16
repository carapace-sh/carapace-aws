package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2Cmd = &cobra.Command{
	Use:   "elbv2",
	Short: "Elastic Load Balancing\n\nA load balancer distributes incoming traffic across targets, such as your EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2Cmd).Standalone()

	rootCmd.AddCommand(elbv2Cmd)
}
