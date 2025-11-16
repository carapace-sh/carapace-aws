package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_createGraphUsingImportTaskCmd = &cobra.Command{
	Use:   "create-graph-using-import-task",
	Short: "Creates a new Neptune Analytics graph and imports data into it, either from Amazon Simple Storage Service (S3) or from a Neptune database or a Neptune database snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_createGraphUsingImportTaskCmd).Standalone()

	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("blank-node-handling", "", "The method to handle blank nodes in the dataset.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("deletion-protection", false, "Indicates whether or not to enable deletion protection on the graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("fail-on-error", false, "If set to `true`, the task halts when an import error is encountered.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("format", "", "Specifies the format of S3 data to be imported.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("graph-name", "", "A name for the new Neptune Analytics graph to be created.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("import-options", "", "Contains options for controlling the import process.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("kms-key-identifier", "", "Specifies a KMS key to use to encrypt data imported into the new graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("max-provisioned-memory", "", "The maximum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("min-provisioned-memory", "", "The minimum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("no-deletion-protection", false, "Indicates whether or not to enable deletion protection on the graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("no-fail-on-error", false, "If set to `true`, the task halts when an import error is encountered.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("no-public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("parquet-type", "", "The parquet type of the import task.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().Bool("public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("replica-count", "", "The number of replicas in other AZs to provision on the new graph after import.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("role-arn", "", "The ARN of the IAM role that will allow access to the data that is to be imported.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("source", "", "A URL identifying to the location of the data to be imported.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("tags", "", "Adds metadata tags to the new graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.Flags().String("vector-search-configuration", "", "Specifies the number of dimensions for vector embeddings that will be loaded into the graph.")
	neptuneGraph_createGraphUsingImportTaskCmd.MarkFlagRequired("graph-name")
	neptuneGraph_createGraphUsingImportTaskCmd.Flag("no-deletion-protection").Hidden = true
	neptuneGraph_createGraphUsingImportTaskCmd.Flag("no-fail-on-error").Hidden = true
	neptuneGraph_createGraphUsingImportTaskCmd.Flag("no-public-connectivity").Hidden = true
	neptuneGraph_createGraphUsingImportTaskCmd.MarkFlagRequired("role-arn")
	neptuneGraph_createGraphUsingImportTaskCmd.MarkFlagRequired("source")
	neptuneGraphCmd.AddCommand(neptuneGraph_createGraphUsingImportTaskCmd)
}
