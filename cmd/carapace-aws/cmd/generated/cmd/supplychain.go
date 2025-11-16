package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychainCmd = &cobra.Command{
	Use:   "supplychain",
	Short: "AWS Supply Chain is a cloud-based application that works with your enterprise resource planning (ERP) and supply chain management systems.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychainCmd).Standalone()

	rootCmd.AddCommand(supplychainCmd)
}
