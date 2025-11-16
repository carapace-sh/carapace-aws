package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listVpcConnectorsCmd = &cobra.Command{
	Use:   "list-vpc-connectors",
	Short: "Returns a list of App Runner VPC connectors in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listVpcConnectorsCmd).Standalone()

	apprunner_listVpcConnectorsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
	apprunner_listVpcConnectorsCmd.Flags().String("next-token", "", "A token from a previous result page.")
	apprunnerCmd.AddCommand(apprunner_listVpcConnectorsCmd)
}
