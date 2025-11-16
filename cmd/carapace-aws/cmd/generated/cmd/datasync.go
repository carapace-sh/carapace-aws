package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasyncCmd = &cobra.Command{
	Use:   "datasync",
	Short: "DataSync\n\nDataSync is an online data movement service that simplifies data migration and helps you quickly, easily, and securely transfer your file or object data to, from, and between Amazon Web Services storage services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasyncCmd).Standalone()

	rootCmd.AddCommand(datasyncCmd)
}
