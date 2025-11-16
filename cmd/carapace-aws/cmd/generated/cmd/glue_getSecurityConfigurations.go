package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSecurityConfigurationsCmd = &cobra.Command{
	Use:   "get-security-configurations",
	Short: "Retrieves a list of all security configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSecurityConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSecurityConfigurationsCmd).Standalone()

		glue_getSecurityConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		glue_getSecurityConfigurationsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_getSecurityConfigurationsCmd)
}
