package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteLakeFormationIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "delete-lake-formation-identity-center-configuration",
	Short: "Deletes an IAM Identity Center connection with Lake Formation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteLakeFormationIdentityCenterConfigurationCmd).Standalone()

	lakeformation_deleteLakeFormationIdentityCenterConfigurationCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformationCmd.AddCommand(lakeformation_deleteLakeFormationIdentityCenterConfigurationCmd)
}
