package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createGlueIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "create-glue-identity-center-configuration",
	Short: "Creates a new Glue Identity Center configuration to enable integration between Glue and Amazon Web Services IAM Identity Center for authentication and authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createGlueIdentityCenterConfigurationCmd).Standalone()

	glue_createGlueIdentityCenterConfigurationCmd.Flags().String("instance-arn", "", "The Amazon Resource Name (ARN) of the Identity Center instance to be associated with the Glue configuration.")
	glue_createGlueIdentityCenterConfigurationCmd.Flags().String("scopes", "", "A list of Identity Center scopes that define the permissions and access levels for the Glue configuration.")
	glue_createGlueIdentityCenterConfigurationCmd.Flags().String("user-background-sessions-enabled", "", "Specifies whether users can run background sessions when using Identity Center authentication with Glue services.")
	glue_createGlueIdentityCenterConfigurationCmd.MarkFlagRequired("instance-arn")
	glueCmd.AddCommand(glue_createGlueIdentityCenterConfigurationCmd)
}
