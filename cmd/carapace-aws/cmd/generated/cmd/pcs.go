package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcsCmd = &cobra.Command{
	Use:   "pcs",
	Short: "Parallel Computing Service (PCS) is a managed service that makes it easier for you to run and scale your high performance computing (HPC) workloads, and build scientific and engineering models on Amazon Web Services using Slurm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcsCmd).Standalone()

	rootCmd.AddCommand(pcsCmd)
}
