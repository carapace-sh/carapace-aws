package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists all applications associated with the instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listApplicationsCmd).Standalone()

		ssoAdmin_listApplicationsCmd.Flags().String("filter", "", "Filters response results.")
		ssoAdmin_listApplicationsCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center application under which the operation will run.")
		ssoAdmin_listApplicationsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		ssoAdmin_listApplicationsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ssoAdmin_listApplicationsCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationsCmd)
}
