package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getComponentCmd = &cobra.Command{
	Use:   "get-component",
	Short: "Get detailed data for a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getComponentCmd).Standalone()

	proton_getComponentCmd.Flags().String("name", "", "The name of the component that you want to get the detailed data for.")
	proton_getComponentCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_getComponentCmd)
}
