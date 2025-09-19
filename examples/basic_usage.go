package main

import (
	"fmt"
	"log"

	"github.com/draper1/esent-go/api"
	"github.com/draper1/esent-go/types"
)

func main() {
	fmt.Println("ESENT Go - Basic Usage Example")
	fmt.Println("===============================")

	// Create a new ESENT API instance
	jetApi := api.NewJetApi()
	if jetApi == nil {
		log.Fatal("Failed to create JetApi instance")
	}

	// Create an ESENT instance
	fmt.Println("Creating ESENT instance...")
	instance, err := jetApi.JetCreateInstance2("ExampleInstance", "Example ESENT Instance", types.CreateInstanceGrbitNone)
	if err != nil {
		log.Fatalf("Failed to create ESENT instance: %v", err)
	}
	fmt.Printf("Created instance: %+v\n", instance)

	// Initialize the instance
	fmt.Println("Initializing ESENT instance...")
	err = jetApi.JetInit2(instance, types.InitGrbitNone)
	if err != nil {
		log.Fatalf("Failed to initialize ESENT instance: %v", err)
	}
	fmt.Println("Instance initialized successfully")

	// Begin a session
	fmt.Println("Beginning session...")
	sesid := &types.JET_SESID{}
	err = jetApi.JetBeginSession(instance, sesid, "", "")
	if err != nil {
		log.Fatalf("Failed to begin session: %v", err)
	}
	fmt.Printf("Session started: %+v\n", sesid)

	// Create a database
	fmt.Println("Creating database...")
	dbid, err := jetApi.JetCreateDatabase(sesid, "example.db", "example.db", types.CreateDatabaseGrbitNone)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	fmt.Printf("Database created: %+v\n", dbid)

	// Create a table
	fmt.Println("Creating table...")
	tableid, err := jetApi.JetCreateTable(sesid, dbid, "ExampleTable", 0, 0)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Printf("Table created: %+v\n", tableid)

	// Create a column
	fmt.Println("Creating column...")
	columnid := &types.JET_COLUMNID{Value: 0x12345678}
	columndef := &types.JET_COLUMNDEF{
		Columnid:   *columnid,
		Coltyp:     types.ColtypText,
		Grbit:      0,
		CbMax:      255,
		Cp:         0,
		ColumnName: "ExampleColumn",
	}

	createdColumnid, err := jetApi.JetAddColumn(sesid, tableid, "ExampleColumn", columndef, nil, 0)
	if err != nil {
		log.Fatalf("Failed to add column: %v", err)
	}
	fmt.Printf("Column created: %+v\n", createdColumnid)

	// Clean up
	fmt.Println("Cleaning up...")

	err = jetApi.JetCloseTable(sesid, tableid)
	if err != nil {
		log.Printf("Warning: Failed to close table: %v", err)
	}

	err = jetApi.JetCloseDatabase(sesid, dbid, 0)
	if err != nil {
		log.Printf("Warning: Failed to close database: %v", err)
	}

	err = jetApi.JetEndSession(sesid, 0)
	if err != nil {
		log.Printf("Warning: Failed to end session: %v", err)
	}

	err = jetApi.JetTerm2(instance, types.TermGrbitNone)
	if err != nil {
		log.Printf("Warning: Failed to terminate instance: %v", err)
	}

	fmt.Println("Cleanup completed")
	fmt.Println("Example completed successfully!")
}
