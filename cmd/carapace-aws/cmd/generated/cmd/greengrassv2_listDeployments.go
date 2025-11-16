package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Retrieves a paginated list of deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_listDeploymentsCmd).Standalone()

		greengrassv2_listDeploymentsCmd.Flags().String("history-filter", "", "The filter for the list of deployments.")
		greengrassv2_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
		greengrassv2_listDeploymentsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		greengrassv2_listDeploymentsCmd.Flags().String("parent-target-arn", "", "The parent deployment's target [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) within a subdeployment.")
		greengrassv2_listDeploymentsCmd.Flags().String("target-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the target IoT thing or thing group.")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_listDeploymentsCmd)
}
