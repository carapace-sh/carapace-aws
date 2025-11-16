package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsCmd = &cobra.Command{
	Use:   "ds",
	Short: "Directory Service\n\nDirectory Service is a web service that makes it easy for you to setup and run directories in the Amazon Web Services cloud, or connect your Amazon Web Services resources with an existing self-managed Microsoft Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsCmd).Standalone()

	rootCmd.AddCommand(dsCmd)
}
