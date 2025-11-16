package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteResourceCmd = &cobra.Command{
	Use:   "delete-resource",
	Short: "Deletes the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteResourceCmd).Standalone()

	workmail_deleteResourceCmd.Flags().String("organization-id", "", "The identifier associated with the organization from which the resource is deleted.")
	workmail_deleteResourceCmd.Flags().String("resource-id", "", "The identifier of the resource to be deleted.")
	workmail_deleteResourceCmd.MarkFlagRequired("organization-id")
	workmail_deleteResourceCmd.MarkFlagRequired("resource-id")
	workmailCmd.AddCommand(workmail_deleteResourceCmd)
}
