package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getEnabledStandardsCmd = &cobra.Command{
	Use:   "get-enabled-standards",
	Short: "Returns a list of the standards that are currently enabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getEnabledStandardsCmd).Standalone()

	securityhub_getEnabledStandardsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	securityhub_getEnabledStandardsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhub_getEnabledStandardsCmd.Flags().String("standards-subscription-arns", "", "The list of the standards subscription ARNs for the standards to retrieve.")
	securityhubCmd.AddCommand(securityhub_getEnabledStandardsCmd)
}
