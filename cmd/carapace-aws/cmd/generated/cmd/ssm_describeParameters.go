package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeParametersCmd = &cobra.Command{
	Use:   "describe-parameters",
	Short: "Lists the parameters in your Amazon Web Services account or the parameters shared with you when you enable the [Shared](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribeParameters.html#systemsmanager-DescribeParameters-request-Shared) option.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeParametersCmd).Standalone()

	ssm_describeParametersCmd.Flags().String("filters", "", "This data type is deprecated.")
	ssm_describeParametersCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeParametersCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeParametersCmd.Flags().Bool("no-shared", false, "Lists parameters that are shared with you.")
	ssm_describeParametersCmd.Flags().String("parameter-filters", "", "Filters to limit the request results.")
	ssm_describeParametersCmd.Flags().Bool("shared", false, "Lists parameters that are shared with you.")
	ssm_describeParametersCmd.Flag("no-shared").Hidden = true
	ssmCmd.AddCommand(ssm_describeParametersCmd)
}
