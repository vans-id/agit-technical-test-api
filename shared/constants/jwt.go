package constants

import (
	"time"
)

var (
	JWT_ISS = "agit-technical-test-api"
	JWT_IAT = time.Now().Unix()
	JWT_EXP = time.Now().Add(time.Hour * 24).Unix()
)
