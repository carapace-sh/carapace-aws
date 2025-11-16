package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolverCmd = &cobra.Command{
	Use:   "route53resolver",
	Short: "When you create a VPC using Amazon VPC, you automatically get DNS resolution within the VPC from Route 53 Resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolverCmd).Standalone()

	})
	rootCmd.AddCommand(route53resolverCmd)
}
