package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateNumberOfDomainControllersCmd = &cobra.Command{
	Use:   "update-number-of-domain-controllers",
	Short: "Adds or removes domain controllers to or from the directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateNumberOfDomainControllersCmd).Standalone()

	ds_updateNumberOfDomainControllersCmd.Flags().String("desired-number", "", "The number of domain controllers desired in the directory.")
	ds_updateNumberOfDomainControllersCmd.Flags().String("directory-id", "", "Identifier of the directory to which the domain controllers will be added or removed.")
	ds_updateNumberOfDomainControllersCmd.MarkFlagRequired("desired-number")
	ds_updateNumberOfDomainControllersCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_updateNumberOfDomainControllersCmd)
}
