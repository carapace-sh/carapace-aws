package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createAssociationBatchCmd = &cobra.Command{
	Use:   "create-association-batch",
	Short: "Associates the specified Amazon Web Services Systems Manager document (SSM document) with the specified managed nodes or targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createAssociationBatchCmd).Standalone()

	ssm_createAssociationBatchCmd.Flags().String("entries", "", "One or more associations.")
	ssm_createAssociationBatchCmd.MarkFlagRequired("entries")
	ssmCmd.AddCommand(ssm_createAssociationBatchCmd)
}
