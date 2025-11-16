package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listSecurityControlDefinitionsCmd = &cobra.Command{
	Use:   "list-security-control-definitions",
	Short: "Lists all of the security controls that apply to a specified standard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listSecurityControlDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listSecurityControlDefinitionsCmd).Standalone()

		securityhub_listSecurityControlDefinitionsCmd.Flags().String("max-results", "", "An optional parameter that limits the total results of the API response to the specified number.")
		securityhub_listSecurityControlDefinitionsCmd.Flags().String("next-token", "", "Optional pagination parameter.")
		securityhub_listSecurityControlDefinitionsCmd.Flags().String("standards-arn", "", "The Amazon Resource Name (ARN) of the standard that you want to view controls for.")
	})
	securityhubCmd.AddCommand(securityhub_listSecurityControlDefinitionsCmd)
}
