package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAccessPoliciesCmd = &cobra.Command{
	Use:   "list-access-policies",
	Short: "Retrieves a paginated list of access policies for an identity (an IAM Identity Center user, an IAM Identity Center group, or an IAM user) or an IoT SiteWise Monitor resource (a portal or project).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAccessPoliciesCmd).Standalone()

	iotsitewise_listAccessPoliciesCmd.Flags().String("iam-arn", "", "The ARN of the IAM user.")
	iotsitewise_listAccessPoliciesCmd.Flags().String("identity-id", "", "The ID of the identity.")
	iotsitewise_listAccessPoliciesCmd.Flags().String("identity-type", "", "The type of identity (IAM Identity Center user, IAM Identity Center group, or IAM user).")
	iotsitewise_listAccessPoliciesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listAccessPoliciesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listAccessPoliciesCmd.Flags().String("resource-id", "", "The ID of the resource.")
	iotsitewise_listAccessPoliciesCmd.Flags().String("resource-type", "", "The type of resource (portal or project).")
	iotsitewiseCmd.AddCommand(iotsitewise_listAccessPoliciesCmd)
}
