package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Set up a cluster configuration for connecting to the master node (write operations)
	masterCluster := gocql.NewCluster("localhost")
	masterCluster.Keyspace = "test_keyspace"
	masterCluster.Consistency = gocql.Quorum // Use QUORUM or other consistency levels as needed
	masterCluster.ProtoVersion = 4           // Adjust based on your Cassandra version

	// Connect to the master node
	masterSession, err := masterCluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to master node: %v", err)
	}
	defer masterSession.Close()

	// Write data to the master node
	writeQuery := "INSERT INTO test_table (id, name) VALUES (3, 'mesut')"
	if err := masterSession.Query(writeQuery, gocql.TimeUUID(), "Master Write").Exec(); err != nil {
		log.Fatalf("Failed to write data to master node: %v", err)
	}
	fmt.Println("Data written to master node.")

	// Set up a cluster configuration for connecting to the slave node (read operations)
	slaveCluster := gocql.NewCluster("localhost:9043")
	slaveCluster.Keyspace = "test_keyspace"
	slaveCluster.Consistency = gocql.One // Use ONE consistency level for read-only operations
	slaveCluster.ProtoVersion = 4        // Adjust based on your Cassandra version

	// Connect to the slave node
	slaveSession, err := slaveCluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to slave node: %v", err)
	}
	defer slaveSession.Close()

	// Read data from the slave node
	var id gocql.UUID
	var name string
	readQuery := "SELECT id, name FROM test_table LIMIT 3"
	if err := slaveSession.Query(readQuery).Consistency(gocql.One).Scan(&id, &name); err != nil {
		log.Fatalf("Failed to read data from slave node: %v", err)
	}
	fmt.Printf("Read data from slave node: id=%v, name=%s\n", id, name)
}
