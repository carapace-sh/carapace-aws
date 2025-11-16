package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getGlueIdentityCenterConfigurationCmd = &cobra.Command{
	Use:   "get-glue-identity-center-configuration",
	Short: "Retrieves the current Glue Identity Center configuration details, including the associated Identity Center instance and application information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getGlueIdentityCenterConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getGlueIdentityCenterConfigurationCmd).Standalone()

	})
	glueCmd.AddCommand(glue_getGlueIdentityCenterConfigurationCmd)
}
