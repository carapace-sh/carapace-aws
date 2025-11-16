package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getAutoSnapshotsCmd = &cobra.Command{
	Use:   "get-auto-snapshots",
	Short: "Returns the available automatic snapshots for an instance or disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getAutoSnapshotsCmd).Standalone()

	lightsail_getAutoSnapshotsCmd.Flags().String("resource-name", "", "The name of the source instance or disk from which to get automatic snapshot information.")
	lightsail_getAutoSnapshotsCmd.MarkFlagRequired("resource-name")
	lightsailCmd.AddCommand(lightsail_getAutoSnapshotsCmd)
}
