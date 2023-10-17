package Gobase

import "GOdeau/godeau/Gobase/GoDriver"

func Init(user string, password string, host string, database string) (*Database, error) {
	return Connect(GoDriver.POSTGRES, user, password, host, database)
}
