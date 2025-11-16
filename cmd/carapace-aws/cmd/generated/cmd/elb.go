package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbCmd = &cobra.Command{
	Use:   "elb",
	Short: "Elastic Load Balancing\n\nA load balancer can distribute incoming traffic across your EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbCmd).Standalone()

	})
	rootCmd.AddCommand(elbCmd)
}
