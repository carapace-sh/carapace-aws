package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateResourceCmd = &cobra.Command{
	Use:   "update-resource",
	Short: "Updates the data access role used for vending access to the given (registered) resource in Lake Formation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_updateResourceCmd).Standalone()

		lakeformation_updateResourceCmd.Flags().String("hybrid-access-enabled", "", "Specifies whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.")
		lakeformation_updateResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		lakeformation_updateResourceCmd.Flags().String("role-arn", "", "The new role to use for the given resource registered in Lake Formation.")
		lakeformation_updateResourceCmd.Flags().String("with-federation", "", "Whether or not the resource is a federated resource.")
		lakeformation_updateResourceCmd.MarkFlagRequired("resource-arn")
		lakeformation_updateResourceCmd.MarkFlagRequired("role-arn")
	})
	lakeformationCmd.AddCommand(lakeformation_updateResourceCmd)
}
