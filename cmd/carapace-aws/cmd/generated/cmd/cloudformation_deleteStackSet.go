package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deleteStackSetCmd = &cobra.Command{
	Use:   "delete-stack-set",
	Short: "Deletes a StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deleteStackSetCmd).Standalone()

	cloudformation_deleteStackSetCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_deleteStackSetCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you're deleting.")
	cloudformation_deleteStackSetCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_deleteStackSetCmd)
}
