package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeResourceGroupsCmd = &cobra.Command{
	Use:   "describe-resource-groups",
	Short: "Describes the resource groups that are specified by the ARNs of the resource groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeResourceGroupsCmd).Standalone()

	inspector_describeResourceGroupsCmd.Flags().String("resource-group-arns", "", "The ARN that specifies the resource group that you want to describe.")
	inspector_describeResourceGroupsCmd.MarkFlagRequired("resource-group-arns")
	inspectorCmd.AddCommand(inspector_describeResourceGroupsCmd)
}
