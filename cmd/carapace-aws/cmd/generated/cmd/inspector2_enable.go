package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables Amazon Inspector scans for one or more Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_enableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_enableCmd).Standalone()

		inspector2_enableCmd.Flags().String("account-ids", "", "A list of account IDs you want to enable Amazon Inspector scans for.")
		inspector2_enableCmd.Flags().String("client-token", "", "The idempotency token for the request.")
		inspector2_enableCmd.Flags().String("resource-types", "", "The resource scan types you want to enable.")
		inspector2_enableCmd.MarkFlagRequired("resource-types")
	})
	inspector2Cmd.AddCommand(inspector2_enableCmd)
}
