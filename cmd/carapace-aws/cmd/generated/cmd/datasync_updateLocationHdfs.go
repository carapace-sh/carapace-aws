package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationHdfsCmd = &cobra.Command{
	Use:   "update-location-hdfs",
	Short: "Modifies the following configuration parameters of the Hadoop Distributed File System (HDFS) transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationHdfsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_updateLocationHdfsCmd).Standalone()

		datasync_updateLocationHdfsCmd.Flags().String("agent-arns", "", "The Amazon Resource Names (ARNs) of the DataSync agents that can connect to your HDFS cluster.")
		datasync_updateLocationHdfsCmd.Flags().String("authentication-type", "", "The type of authentication used to determine the identity of the user.")
		datasync_updateLocationHdfsCmd.Flags().String("block-size", "", "The size of the data blocks to write into the HDFS cluster.")
		datasync_updateLocationHdfsCmd.Flags().String("kerberos-keytab", "", "The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys.")
		datasync_updateLocationHdfsCmd.Flags().String("kerberos-krb5-conf", "", "The `krb5.conf` file that contains the Kerberos configuration information.")
		datasync_updateLocationHdfsCmd.Flags().String("kerberos-principal", "", "The Kerberos principal with access to the files and folders on the HDFS cluster.")
		datasync_updateLocationHdfsCmd.Flags().String("kms-key-provider-uri", "", "The URI of the HDFS cluster's Key Management Server (KMS).")
		datasync_updateLocationHdfsCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the source HDFS cluster location.")
		datasync_updateLocationHdfsCmd.Flags().String("name-nodes", "", "The NameNode that manages the HDFS namespace.")
		datasync_updateLocationHdfsCmd.Flags().String("qop-configuration", "", "The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer privacy settings configured on the Hadoop Distributed File System (HDFS) cluster.")
		datasync_updateLocationHdfsCmd.Flags().String("replication-factor", "", "The number of DataNodes to replicate the data to when writing to the HDFS cluster.")
		datasync_updateLocationHdfsCmd.Flags().String("simple-user", "", "The user name used to identify the client on the host operating system.")
		datasync_updateLocationHdfsCmd.Flags().String("subdirectory", "", "A subdirectory in the HDFS cluster.")
		datasync_updateLocationHdfsCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_updateLocationHdfsCmd)
}
