package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteLicenseEndpointCmd = &cobra.Command{
	Use:   "delete-license-endpoint",
	Short: "Deletes a license endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteLicenseEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_deleteLicenseEndpointCmd).Standalone()

		deadline_deleteLicenseEndpointCmd.Flags().String("license-endpoint-id", "", "The license endpoint ID of the license endpoint to delete.")
		deadline_deleteLicenseEndpointCmd.MarkFlagRequired("license-endpoint-id")
	})
	deadlineCmd.AddCommand(deadline_deleteLicenseEndpointCmd)
}
