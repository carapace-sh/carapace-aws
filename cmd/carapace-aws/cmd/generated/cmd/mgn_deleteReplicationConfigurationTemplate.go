package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "delete-replication-configuration-template",
	Short: "Deletes a single Replication Configuration Template by ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteReplicationConfigurationTemplateCmd).Standalone()

	mgn_deleteReplicationConfigurationTemplateCmd.Flags().String("replication-configuration-template-id", "", "Request to delete Replication Configuration Template from service by Replication Configuration Template ID.")
	mgn_deleteReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-configuration-template-id")
	mgnCmd.AddCommand(mgn_deleteReplicationConfigurationTemplateCmd)
}
