package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createResourceCmd = &cobra.Command{
	Use:   "create-resource",
	Short: "Creates a new WorkMail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createResourceCmd).Standalone()

	workmail_createResourceCmd.Flags().String("description", "", "Resource description.")
	workmail_createResourceCmd.Flags().Bool("hidden-from-global-address-list", false, "If this parameter is enabled, the resource will be hidden from the address book.")
	workmail_createResourceCmd.Flags().String("name", "", "The name of the new resource.")
	workmail_createResourceCmd.Flags().Bool("no-hidden-from-global-address-list", false, "If this parameter is enabled, the resource will be hidden from the address book.")
	workmail_createResourceCmd.Flags().String("organization-id", "", "The identifier associated with the organization for which the resource is created.")
	workmail_createResourceCmd.Flags().String("type", "", "The type of the new resource.")
	workmail_createResourceCmd.MarkFlagRequired("name")
	workmail_createResourceCmd.Flag("no-hidden-from-global-address-list").Hidden = true
	workmail_createResourceCmd.MarkFlagRequired("organization-id")
	workmail_createResourceCmd.MarkFlagRequired("type")
	workmailCmd.AddCommand(workmail_createResourceCmd)
}
