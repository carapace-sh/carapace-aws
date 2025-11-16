package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateGlueIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "update-glue-identity-center-configuration",
	Short: "Updates the existing Glue Identity Center configuration, allowing modification of scopes and permissions for the integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateGlueIdentityCenterConfigurationCmd).Standalone()

	glue_updateGlueIdentityCenterConfigurationCmd.Flags().String("scopes", "", "A list of Identity Center scopes that define the updated permissions and access levels for the Glue configuration.")
	glue_updateGlueIdentityCenterConfigurationCmd.Flags().String("user-background-sessions-enabled", "", "Specifies whether users can run background sessions when using Identity Center authentication with Glue services.")
	glueCmd.AddCommand(glue_updateGlueIdentityCenterConfigurationCmd)
}
