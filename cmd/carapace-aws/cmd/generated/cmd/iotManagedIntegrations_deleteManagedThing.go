package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteManagedThingCmd = &cobra.Command{
	Use:   "delete-managed-thing",
	Short: "Delete a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteManagedThingCmd).Standalone()

	iotManagedIntegrations_deleteManagedThingCmd.Flags().Bool("force", false, "When set to `TRUE`, a forceful deteletion of the managed thing will occur.")
	iotManagedIntegrations_deleteManagedThingCmd.Flags().String("identifier", "", "The id of the managed thing.")
	iotManagedIntegrations_deleteManagedThingCmd.Flags().Bool("no-force", false, "When set to `TRUE`, a forceful deteletion of the managed thing will occur.")
	iotManagedIntegrations_deleteManagedThingCmd.MarkFlagRequired("identifier")
	iotManagedIntegrations_deleteManagedThingCmd.Flag("no-force").Hidden = true
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteManagedThingCmd)
}
