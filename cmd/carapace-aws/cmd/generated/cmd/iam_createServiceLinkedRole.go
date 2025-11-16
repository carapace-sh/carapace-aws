package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createServiceLinkedRoleCmd = &cobra.Command{
	Use:   "create-service-linked-role",
	Short: "Creates an IAM role that is linked to a specific Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createServiceLinkedRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createServiceLinkedRoleCmd).Standalone()

		iam_createServiceLinkedRoleCmd.Flags().String("awsservice-name", "", "The service principal for the Amazon Web Services service to which this role is attached.")
		iam_createServiceLinkedRoleCmd.Flags().String("custom-suffix", "", "A string that you provide, which is combined with the service-provided prefix to form the complete role name.")
		iam_createServiceLinkedRoleCmd.Flags().String("description", "", "The description of the role.")
		iam_createServiceLinkedRoleCmd.MarkFlagRequired("awsservice-name")
	})
	iamCmd.AddCommand(iam_createServiceLinkedRoleCmd)
}
