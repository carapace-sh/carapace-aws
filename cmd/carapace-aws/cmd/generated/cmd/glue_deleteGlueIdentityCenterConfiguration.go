package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteGlueIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "delete-glue-identity-center-configuration",
	Short: "Deletes the existing Glue Identity Center configuration, removing the integration between Glue and Amazon Web Services IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteGlueIdentityCenterConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteGlueIdentityCenterConfigurationCmd).Standalone()

	})
	glueCmd.AddCommand(glue_deleteGlueIdentityCenterConfigurationCmd)
}
