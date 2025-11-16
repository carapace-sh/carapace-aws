package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateAssociationStatusCmd = &cobra.Command{
	Use:   "update-association-status",
	Short: "Updates the status of the Amazon Web Services Systems Manager document (SSM document) associated with the specified managed node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateAssociationStatusCmd).Standalone()

	ssm_updateAssociationStatusCmd.Flags().String("association-status", "", "The association status.")
	ssm_updateAssociationStatusCmd.Flags().String("instance-id", "", "The managed node ID.")
	ssm_updateAssociationStatusCmd.Flags().String("name", "", "The name of the SSM document.")
	ssm_updateAssociationStatusCmd.MarkFlagRequired("association-status")
	ssm_updateAssociationStatusCmd.MarkFlagRequired("instance-id")
	ssm_updateAssociationStatusCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_updateAssociationStatusCmd)
}
