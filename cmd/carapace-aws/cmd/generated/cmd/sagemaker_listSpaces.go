package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listSpacesCmd = &cobra.Command{
	Use:   "list-spaces",
	Short: "Lists spaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listSpacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listSpacesCmd).Standalone()

		sagemaker_listSpacesCmd.Flags().String("domain-id-equals", "", "A parameter to search for the domain ID.")
		sagemaker_listSpacesCmd.Flags().String("max-results", "", "This parameter defines the maximum number of results that can be return in a single response.")
		sagemaker_listSpacesCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
		sagemaker_listSpacesCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemaker_listSpacesCmd.Flags().String("sort-order", "", "The sort order for the results.")
		sagemaker_listSpacesCmd.Flags().String("space-name-contains", "", "A parameter by which to filter the results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listSpacesCmd)
}
