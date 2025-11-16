package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeInputCmd = &cobra.Command{
	Use:   "describe-input",
	Short: "Produces details about an input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeInputCmd).Standalone()

	medialive_describeInputCmd.Flags().String("input-id", "", "Unique ID of the input")
	medialive_describeInputCmd.MarkFlagRequired("input-id")
	medialiveCmd.AddCommand(medialive_describeInputCmd)
}
