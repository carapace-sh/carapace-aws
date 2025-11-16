package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getLicenseEndpointCmd = &cobra.Command{
	Use:   "get-license-endpoint",
	Short: "Gets a licence endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getLicenseEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getLicenseEndpointCmd).Standalone()

		deadline_getLicenseEndpointCmd.Flags().String("license-endpoint-id", "", "The license endpoint ID.")
		deadline_getLicenseEndpointCmd.MarkFlagRequired("license-endpoint-id")
	})
	deadlineCmd.AddCommand(deadline_getLicenseEndpointCmd)
}
