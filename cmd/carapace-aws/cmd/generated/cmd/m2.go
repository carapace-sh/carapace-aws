package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2Cmd = &cobra.Command{
	Use:   "m2",
	Short: "Amazon Web Services Mainframe Modernization provides tools and resources to help you plan and implement migration and modernization from mainframes to Amazon Web Services managed runtime environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2Cmd).Standalone()

	rootCmd.AddCommand(m2Cmd)
}
