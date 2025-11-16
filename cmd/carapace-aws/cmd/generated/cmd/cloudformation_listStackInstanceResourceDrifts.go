package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackInstanceResourceDriftsCmd = &cobra.Command{
	Use:   "list-stack-instance-resource-drifts",
	Short: "Returns drift information for resources in a stack instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackInstanceResourceDriftsCmd).Standalone()

	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("operation-id", "", "The unique ID of the drift operation.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("stack-instance-account", "", "The name of the Amazon Web Services account that you want to list resource drifts for.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("stack-instance-region", "", "The name of the Region where you want to list resource drifts.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("stack-instance-resource-drift-statuses", "", "The resource drift status of the stack instance.")
	cloudformation_listStackInstanceResourceDriftsCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to list drifted resources for.")
	cloudformation_listStackInstanceResourceDriftsCmd.MarkFlagRequired("operation-id")
	cloudformation_listStackInstanceResourceDriftsCmd.MarkFlagRequired("stack-instance-account")
	cloudformation_listStackInstanceResourceDriftsCmd.MarkFlagRequired("stack-instance-region")
	cloudformation_listStackInstanceResourceDriftsCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_listStackInstanceResourceDriftsCmd)
}
