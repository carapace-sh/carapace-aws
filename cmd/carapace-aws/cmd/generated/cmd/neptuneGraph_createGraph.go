package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_createGraphCmd = &cobra.Command{
	Use:   "create-graph",
	Short: "Creates a new Neptune Analytics graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_createGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_createGraphCmd).Standalone()

		neptuneGraph_createGraphCmd.Flags().Bool("deletion-protection", false, "Indicates whether or not to enable deletion protection on the graph.")
		neptuneGraph_createGraphCmd.Flags().String("graph-name", "", "A name for the new Neptune Analytics graph to be created.")
		neptuneGraph_createGraphCmd.Flags().String("kms-key-identifier", "", "Specifies a KMS key to use to encrypt data in the new graph.")
		neptuneGraph_createGraphCmd.Flags().Bool("no-deletion-protection", false, "Indicates whether or not to enable deletion protection on the graph.")
		neptuneGraph_createGraphCmd.Flags().Bool("no-public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
		neptuneGraph_createGraphCmd.Flags().String("provisioned-memory", "", "The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.")
		neptuneGraph_createGraphCmd.Flags().Bool("public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
		neptuneGraph_createGraphCmd.Flags().String("replica-count", "", "The number of replicas in other AZs.")
		neptuneGraph_createGraphCmd.Flags().String("tags", "", "Adds metadata tags to the new graph.")
		neptuneGraph_createGraphCmd.Flags().String("vector-search-configuration", "", "Specifies the number of dimensions for vector embeddings that will be loaded into the graph.")
		neptuneGraph_createGraphCmd.MarkFlagRequired("graph-name")
		neptuneGraph_createGraphCmd.Flag("no-deletion-protection").Hidden = true
		neptuneGraph_createGraphCmd.Flag("no-public-connectivity").Hidden = true
		neptuneGraph_createGraphCmd.MarkFlagRequired("provisioned-memory")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_createGraphCmd)
}
