package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateLakeFormationIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "update-lake-formation-identity-center-configuration",
	Short: "Updates the IAM Identity Center connection parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateLakeFormationIdentityCenterConfigurationCmd).Standalone()

	lakeformation_updateLakeFormationIdentityCenterConfigurationCmd.Flags().String("application-status", "", "Allows to enable or disable the IAM Identity Center connection.")
	lakeformation_updateLakeFormationIdentityCenterConfigurationCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_updateLakeFormationIdentityCenterConfigurationCmd.Flags().String("external-filtering", "", "A list of the account IDs of Amazon Web Services accounts of third-party applications that are allowed to access data managed by Lake Formation.")
	lakeformation_updateLakeFormationIdentityCenterConfigurationCmd.Flags().String("share-recipients", "", "A list of Amazon Web Services account IDs or Amazon Web Services organization/organizational unit ARNs that are allowed to access to access data managed by Lake Formation.")
	lakeformationCmd.AddCommand(lakeformation_updateLakeFormationIdentityCenterConfigurationCmd)
}
