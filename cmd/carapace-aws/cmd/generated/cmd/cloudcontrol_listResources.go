package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "Returns information about the specified resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_listResourcesCmd).Standalone()

	cloudcontrol_listResourcesCmd.Flags().String("max-results", "", "Reserved.")
	cloudcontrol_listResourcesCmd.Flags().String("next-token", "", "If the previous paginated request didn't return all of the remaining results, the response object's `NextToken` parameter value is set to a token.")
	cloudcontrol_listResourcesCmd.Flags().String("resource-model", "", "The resource model to use to select the resources to return.")
	cloudcontrol_listResourcesCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation.")
	cloudcontrol_listResourcesCmd.Flags().String("type-name", "", "The name of the resource type.")
	cloudcontrol_listResourcesCmd.Flags().String("type-version-id", "", "For private resource types, the type version to use in this resource operation.")
	cloudcontrol_listResourcesCmd.MarkFlagRequired("type-name")
	cloudcontrolCmd.AddCommand(cloudcontrol_listResourcesCmd)
}
