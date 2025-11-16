package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteInputCmd = &cobra.Command{
	Use:   "delete-input",
	Short: "Deletes the input end point",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteInputCmd).Standalone()

	medialive_deleteInputCmd.Flags().String("input-id", "", "Unique ID of the input")
	medialive_deleteInputCmd.MarkFlagRequired("input-id")
	medialiveCmd.AddCommand(medialive_deleteInputCmd)
}
