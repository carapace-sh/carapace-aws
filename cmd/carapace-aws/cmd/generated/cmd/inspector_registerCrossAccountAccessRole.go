package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_registerCrossAccountAccessRoleCmd = &cobra.Command{
	Use:   "register-cross-account-access-role",
	Short: "Registers the IAM role that grants Amazon Inspector access to AWS Services needed to perform security assessments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_registerCrossAccountAccessRoleCmd).Standalone()

	inspector_registerCrossAccountAccessRoleCmd.Flags().String("role-arn", "", "The ARN of the IAM role that grants Amazon Inspector access to AWS Services needed to perform security assessments.")
	inspector_registerCrossAccountAccessRoleCmd.MarkFlagRequired("role-arn")
	inspectorCmd.AddCommand(inspector_registerCrossAccountAccessRoleCmd)
}
