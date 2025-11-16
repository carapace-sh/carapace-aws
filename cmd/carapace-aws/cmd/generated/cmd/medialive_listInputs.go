package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listInputsCmd = &cobra.Command{
	Use:   "list-inputs",
	Short: "Produces list of inputs that have been created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listInputsCmd).Standalone()

	medialive_listInputsCmd.Flags().String("max-results", "", "")
	medialive_listInputsCmd.Flags().String("next-token", "", "")
	medialiveCmd.AddCommand(medialive_listInputsCmd)
}
