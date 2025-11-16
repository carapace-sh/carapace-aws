package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putLabelCmd = &cobra.Command{
	Use:   "put-label",
	Short: "Creates or updates label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putLabelCmd).Standalone()

	frauddetector_putLabelCmd.Flags().String("description", "", "The label description.")
	frauddetector_putLabelCmd.Flags().String("name", "", "The label name.")
	frauddetector_putLabelCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_putLabelCmd.MarkFlagRequired("name")
	frauddetectorCmd.AddCommand(frauddetector_putLabelCmd)
}
