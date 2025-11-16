package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydbCmd = &cobra.Command{
	Use:   "memorydb",
	Short: "MemoryDB is a fully managed, Redis OSS-compatible, in-memory database that delivers ultra-fast performance and Multi-AZ durability for modern applications built using microservices architectures.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydbCmd).Standalone()

	rootCmd.AddCommand(memorydbCmd)
}
