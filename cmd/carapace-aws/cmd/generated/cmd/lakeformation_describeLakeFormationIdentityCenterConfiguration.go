package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_describeLakeFormationIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "describe-lake-formation-identity-center-configuration",
	Short: "Retrieves the instance ARN and application ARN for the connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_describeLakeFormationIdentityCenterConfigurationCmd).Standalone()

	lakeformation_describeLakeFormationIdentityCenterConfigurationCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformationCmd.AddCommand(lakeformation_describeLakeFormationIdentityCenterConfigurationCmd)
}
