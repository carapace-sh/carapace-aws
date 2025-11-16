package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describePortalCmd = &cobra.Command{
	Use:   "describe-portal",
	Short: "Retrieves information about a portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describePortalCmd).Standalone()

	iotsitewise_describePortalCmd.Flags().String("portal-id", "", "The ID of the portal.")
	iotsitewise_describePortalCmd.MarkFlagRequired("portal-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describePortalCmd)
}
