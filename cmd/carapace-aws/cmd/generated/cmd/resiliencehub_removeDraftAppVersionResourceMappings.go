package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_removeDraftAppVersionResourceMappingsCmd = &cobra.Command{
	Use:   "remove-draft-app-version-resource-mappings",
	Short: "Removes resource mappings from a draft application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_removeDraftAppVersionResourceMappingsCmd).Standalone()

	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("app-registry-app-names", "", "The names of the registered applications you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("eks-source-names", "", "The names of the Amazon Elastic Kubernetes Service clusters and namespaces you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("logical-stack-names", "", "The names of the CloudFormation stacks you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("resource-group-names", "", "The names of the resource groups you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("resource-names", "", "The names of the resources you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.Flags().String("terraform-source-names", "", "The names of the Terraform sources you want to remove from the resource mappings.")
	resiliencehub_removeDraftAppVersionResourceMappingsCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_removeDraftAppVersionResourceMappingsCmd)
}
