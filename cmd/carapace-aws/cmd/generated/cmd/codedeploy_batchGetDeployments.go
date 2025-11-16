package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetDeploymentsCmd = &cobra.Command{
	Use:   "batch-get-deployments",
	Short: "Gets information about one or more deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_batchGetDeploymentsCmd).Standalone()

		codedeploy_batchGetDeploymentsCmd.Flags().String("deployment-ids", "", "A list of deployment IDs, separated by spaces.")
		codedeploy_batchGetDeploymentsCmd.MarkFlagRequired("deployment-ids")
	})
	codedeployCmd.AddCommand(codedeploy_batchGetDeploymentsCmd)
}
