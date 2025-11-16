package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_batchAssociateUserStackCmd = &cobra.Command{
	Use:   "batch-associate-user-stack",
	Short: "Associates the specified users with the specified stacks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_batchAssociateUserStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_batchAssociateUserStackCmd).Standalone()

		appstream_batchAssociateUserStackCmd.Flags().String("user-stack-associations", "", "The list of UserStackAssociation objects.")
		appstream_batchAssociateUserStackCmd.MarkFlagRequired("user-stack-associations")
	})
	appstreamCmd.AddCommand(appstream_batchAssociateUserStackCmd)
}
