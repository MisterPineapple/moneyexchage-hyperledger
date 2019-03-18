package main

import (
	"errors"
)

// ACL - Access List
type ACL struct {
}

//CreateAccountTX - Access Control Funcion for CreateAccountTX
func (t *ACL) CreateAccountTX(account Account, createdAccount CreateAccountTransaction, txName string) error {
	pass := true
	var reason string

	switch account.Role {
	case "admin":
		return nil

	default:
		pass = false
	}

	if pass {
		return nil
	}

	return errors.New(" " + account.Id + " do not have access to " + txName + ". " + reason)
}