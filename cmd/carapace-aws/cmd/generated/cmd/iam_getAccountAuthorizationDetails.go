package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getAccountAuthorizationDetailsCmd = &cobra.Command{
	Use:   "get-account-authorization-details",
	Short: "Retrieves information about all IAM users, groups, roles, and policies in your Amazon Web Services account, including their relationships to one another.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getAccountAuthorizationDetailsCmd).Standalone()

	iam_getAccountAuthorizationDetailsCmd.Flags().String("filter", "", "A list of entity types used to filter the results.")
	iam_getAccountAuthorizationDetailsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_getAccountAuthorizationDetailsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iamCmd.AddCommand(iam_getAccountAuthorizationDetailsCmd)
}
