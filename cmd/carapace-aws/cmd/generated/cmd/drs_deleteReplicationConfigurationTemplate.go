package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "delete-replication-configuration-template",
	Short: "Deletes a single Replication Configuration Template by ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteReplicationConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_deleteReplicationConfigurationTemplateCmd).Standalone()

		drs_deleteReplicationConfigurationTemplateCmd.Flags().String("replication-configuration-template-id", "", "The ID of the Replication Configuration Template to be deleted.")
		drs_deleteReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-configuration-template-id")
	})
	drsCmd.AddCommand(drs_deleteReplicationConfigurationTemplateCmd)
}
