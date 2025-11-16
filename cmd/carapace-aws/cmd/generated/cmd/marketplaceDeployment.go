package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceDeploymentCmd = &cobra.Command{
	Use:   "marketplace-deployment",
	Short: "The AWS Marketplace Deployment Service supports the Quick Launch experience, which is a deployment option for software as a service (SaaS) products.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceDeploymentCmd).Standalone()

	})
	rootCmd.AddCommand(marketplaceDeploymentCmd)
}
