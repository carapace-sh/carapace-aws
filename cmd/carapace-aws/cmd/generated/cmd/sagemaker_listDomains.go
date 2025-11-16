package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Lists the domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listDomainsCmd).Standalone()

	sagemaker_listDomainsCmd.Flags().String("max-results", "", "This parameter defines the maximum number of results that can be return in a single response.")
	sagemaker_listDomainsCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
	sagemakerCmd.AddCommand(sagemaker_listDomainsCmd)
}
