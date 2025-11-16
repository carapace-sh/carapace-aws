package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteComponentCmd = &cobra.Command{
	Use:   "delete-component",
	Short: "Delete an Proton component resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteComponentCmd).Standalone()

	proton_deleteComponentCmd.Flags().String("name", "", "The name of the component to delete.")
	proton_deleteComponentCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_deleteComponentCmd)
}
