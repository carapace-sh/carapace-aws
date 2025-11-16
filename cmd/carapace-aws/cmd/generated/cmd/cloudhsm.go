package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmCmd = &cobra.Command{
	Use:   "cloudhsm",
	Short: "AWS CloudHSM Service\n\nThis is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmCmd).Standalone()

	rootCmd.AddCommand(cloudhsmCmd)
}
