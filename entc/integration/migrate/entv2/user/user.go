// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "oid"      // FieldAge holds the string denoting the age vertex property in the database.
	FieldAge      = "age"      // FieldName holds the string denoting the name vertex property in the database.
	FieldName     = "name"     // FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname" // FieldPhone holds the string denoting the phone vertex property in the database.
	FieldPhone    = "phone"    // FieldBuffer holds the string denoting the buffer vertex property in the database.
	FieldBuffer   = "buffer"   // FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle    = "title"    // FieldNewName holds the string denoting the new_name vertex property in the database.
	FieldNewName  = "renamed"  // FieldBlob holds the string denoting the blob vertex property in the database.
	FieldBlob     = "blob"     // FieldState holds the string denoting the state vertex property in the database.
	FieldState    = "state"

	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CarTable is the table the holds the car relation/edge.
	CarTable = "cars"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "user_car"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "users"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldNickname,
	FieldPhone,
	FieldBuffer,
	FieldTitle,
	FieldNewName,
	FieldBlob,
	FieldState,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_pets",
}

var (
	// DefaultPhone holds the default value on creation for the phone field.
	DefaultPhone string
	// DefaultTitle holds the default value on creation for the title field.
	DefaultTitle string
)

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
	StateOnline    State = "online"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateLoggedIn, StateLoggedOut, StateOnline:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", s)
	}
}
