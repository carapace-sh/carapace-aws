package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteAssociationCmd = &cobra.Command{
	Use:   "delete-association",
	Short: "Disassociates the specified Amazon Web Services Systems Manager document (SSM document) from the specified managed node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteAssociationCmd).Standalone()

		ssm_deleteAssociationCmd.Flags().String("association-id", "", "The association ID that you want to delete.")
		ssm_deleteAssociationCmd.Flags().String("instance-id", "", "The managed node ID.")
		ssm_deleteAssociationCmd.Flags().String("name", "", "The name of the SSM document.")
	})
	ssmCmd.AddCommand(ssm_deleteAssociationCmd)
}
