package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationHdfsCmd = &cobra.Command{
	Use:   "create-location-hdfs",
	Short: "Creates a transfer *location* for a Hadoop Distributed File System (HDFS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationHdfsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_createLocationHdfsCmd).Standalone()

		datasync_createLocationHdfsCmd.Flags().String("agent-arns", "", "The Amazon Resource Names (ARNs) of the DataSync agents that can connect to your HDFS cluster.")
		datasync_createLocationHdfsCmd.Flags().String("authentication-type", "", "The type of authentication used to determine the identity of the user.")
		datasync_createLocationHdfsCmd.Flags().String("block-size", "", "The size of data blocks to write into the HDFS cluster.")
		datasync_createLocationHdfsCmd.Flags().String("kerberos-keytab", "", "The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys.")
		datasync_createLocationHdfsCmd.Flags().String("kerberos-krb5-conf", "", "The `krb5.conf` file that contains the Kerberos configuration information.")
		datasync_createLocationHdfsCmd.Flags().String("kerberos-principal", "", "The Kerberos principal with access to the files and folders on the HDFS cluster.")
		datasync_createLocationHdfsCmd.Flags().String("kms-key-provider-uri", "", "The URI of the HDFS cluster's Key Management Server (KMS).")
		datasync_createLocationHdfsCmd.Flags().String("name-nodes", "", "The NameNode that manages the HDFS namespace.")
		datasync_createLocationHdfsCmd.Flags().String("qop-configuration", "", "The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster.")
		datasync_createLocationHdfsCmd.Flags().String("replication-factor", "", "The number of DataNodes to replicate the data to when writing to the HDFS cluster.")
		datasync_createLocationHdfsCmd.Flags().String("simple-user", "", "The user name used to identify the client on the host operating system.")
		datasync_createLocationHdfsCmd.Flags().String("subdirectory", "", "A subdirectory in the HDFS cluster.")
		datasync_createLocationHdfsCmd.Flags().String("tags", "", "The key-value pair that represents the tag that you want to add to the location.")
		datasync_createLocationHdfsCmd.MarkFlagRequired("agent-arns")
		datasync_createLocationHdfsCmd.MarkFlagRequired("authentication-type")
		datasync_createLocationHdfsCmd.MarkFlagRequired("name-nodes")
	})
	datasyncCmd.AddCommand(datasync_createLocationHdfsCmd)
}
