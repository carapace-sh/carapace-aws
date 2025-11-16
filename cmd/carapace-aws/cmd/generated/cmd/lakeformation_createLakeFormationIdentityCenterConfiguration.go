package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_createLakeFormationIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "create-lake-formation-identity-center-configuration",
	Short: "Creates an IAM Identity Center connection with Lake Formation to allow IAM Identity Center users and groups to access Data Catalog resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_createLakeFormationIdentityCenterConfigurationCmd).Standalone()

	lakeformation_createLakeFormationIdentityCenterConfigurationCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_createLakeFormationIdentityCenterConfigurationCmd.Flags().String("external-filtering", "", "A list of the account IDs of Amazon Web Services accounts of third-party applications that are allowed to access data managed by Lake Formation.")
	lakeformation_createLakeFormationIdentityCenterConfigurationCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance for which the operation will be executed.")
	lakeformation_createLakeFormationIdentityCenterConfigurationCmd.Flags().String("share-recipients", "", "A list of Amazon Web Services account IDs and/or Amazon Web Services organization/organizational unit ARNs that are allowed to access data managed by Lake Formation.")
	lakeformationCmd.AddCommand(lakeformation_createLakeFormationIdentityCenterConfigurationCmd)
}
