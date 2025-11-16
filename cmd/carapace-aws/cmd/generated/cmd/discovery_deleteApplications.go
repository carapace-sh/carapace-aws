package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_deleteApplicationsCmd = &cobra.Command{
	Use:   "delete-applications",
	Short: "Deletes a list of applications and their associations with configuration items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_deleteApplicationsCmd).Standalone()

	discovery_deleteApplicationsCmd.Flags().String("configuration-ids", "", "Configuration ID of an application to be deleted.")
	discovery_deleteApplicationsCmd.MarkFlagRequired("configuration-ids")
	discoveryCmd.AddCommand(discovery_deleteApplicationsCmd)
}
