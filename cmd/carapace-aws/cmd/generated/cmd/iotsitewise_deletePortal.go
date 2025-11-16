package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deletePortalCmd = &cobra.Command{
	Use:   "delete-portal",
	Short: "Deletes a portal from IoT SiteWise Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deletePortalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deletePortalCmd).Standalone()

		iotsitewise_deletePortalCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deletePortalCmd.Flags().String("portal-id", "", "The ID of the portal to delete.")
		iotsitewise_deletePortalCmd.MarkFlagRequired("portal-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deletePortalCmd)
}
