package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "Creates an alias for a directory and assigns the alias to the directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createAliasCmd).Standalone()

	ds_createAliasCmd.Flags().String("alias", "", "The requested alias.")
	ds_createAliasCmd.Flags().String("directory-id", "", "The identifier of the directory for which to create the alias.")
	ds_createAliasCmd.MarkFlagRequired("alias")
	ds_createAliasCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_createAliasCmd)
}
