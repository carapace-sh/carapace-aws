package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeStandardsCmd = &cobra.Command{
	Use:   "describe-standards",
	Short: "Returns a list of the available standards in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeStandardsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_describeStandardsCmd).Standalone()

		securityhub_describeStandardsCmd.Flags().String("max-results", "", "The maximum number of standards to return.")
		securityhub_describeStandardsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	})
	securityhubCmd.AddCommand(securityhub_describeStandardsCmd)
}
