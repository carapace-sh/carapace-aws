package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_batchDisassociateUserStackCmd = &cobra.Command{
	Use:   "batch-disassociate-user-stack",
	Short: "Disassociates the specified users from the specified stacks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_batchDisassociateUserStackCmd).Standalone()

	appstream_batchDisassociateUserStackCmd.Flags().String("user-stack-associations", "", "The list of UserStackAssociation objects.")
	appstream_batchDisassociateUserStackCmd.MarkFlagRequired("user-stack-associations")
	appstreamCmd.AddCommand(appstream_batchDisassociateUserStackCmd)
}
