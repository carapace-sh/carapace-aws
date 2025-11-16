package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listInstanceProfilesForRoleCmd = &cobra.Command{
	Use:   "list-instance-profiles-for-role",
	Short: "Lists the instance profiles that have the specified associated IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listInstanceProfilesForRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listInstanceProfilesForRoleCmd).Standalone()

		iam_listInstanceProfilesForRoleCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listInstanceProfilesForRoleCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listInstanceProfilesForRoleCmd.Flags().String("role-name", "", "The name of the role to list instance profiles for.")
		iam_listInstanceProfilesForRoleCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_listInstanceProfilesForRoleCmd)
}
