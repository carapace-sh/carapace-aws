package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listVpcIngressConnectionsCmd = &cobra.Command{
	Use:   "list-vpc-ingress-connections",
	Short: "Return a list of App Runner VPC Ingress Connections in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listVpcIngressConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_listVpcIngressConnectionsCmd).Standalone()

		apprunner_listVpcIngressConnectionsCmd.Flags().String("filter", "", "The VPC Ingress Connections to be listed based on either the Service Arn or Vpc Endpoint Id, or both.")
		apprunner_listVpcIngressConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
		apprunner_listVpcIngressConnectionsCmd.Flags().String("next-token", "", "A token from a previous result page.")
	})
	apprunnerCmd.AddCommand(apprunner_listVpcIngressConnectionsCmd)
}
